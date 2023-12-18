package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Translation represents the structure of your document.
type Translation struct {
    ID               primitive.ObjectID `bson:"_id"`
    TranslatorName   string             `bson:"translatorName"`
    Language         string             `bson:"language"`
    Content          string             `bson:"content"`
    ContentEmbedding []float64          `bson:"contentEmbedding"`
    VerseNumber      int                `bson:"verseNumber"`
    ChapterNumber    int                `bson:"chapterNumber"`
}

// getVectorEmbedding simulates the function to compute the vector embedding.
// Replace this with your actual implementation.
func getVectorEmbedding(content string)( []float32, error){
    ctx := context.Background()
// Access your API key as an environment variable (see "Set up your API key" above)
client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_SECRET")))
if err != nil {
    return nil, err
}
defer client.Close()
em := client.EmbeddingModel("embedding-001")
res, err := em.EmbedContent(ctx, genai.Text("The quick brown fox jumps over the lazy dog."))

if err != nil {
    return nil,err
}
    return res.Embedding.Values, nil
}

// updateAllContentEmbeddings updates the contentEmbedding of all Translation documents in the collection.
func updateAllContentEmbeddings(client *mongo.Client) error {
    ctx := context.TODO()
    collection := client.Database("qurandb").Collection("translations")

    cursor, err := collection.Find(ctx, bson.M{"contentEmbedding": []float32{}})
    if err != nil {
        return err
    }
    defer cursor.Close(ctx)

    var totalProcessed, totalFailures int

    for cursor.Next(ctx) {
        totalProcessed++
        var translation Translation
        err := cursor.Decode(&translation)
        if err != nil {
            log.Println("Error decoding document: ", err)
            totalFailures++
            continue
        }

        embedding , err := getVectorEmbedding(translation.Content)
        if err != nil {
            log.Println("Error generating embeddings: ", err)
            totalFailures++
            continue
        }

        filter := bson.M{"_id": translation.ID}
        update := bson.M{"$set": bson.M{"contentEmbedding": embedding}}
        _, err = collection.UpdateOne(ctx, filter, update)
        if err != nil {
            log.Println("Error updating document: ", err)
            totalFailures++
            continue
        }

        // Print progress
        log.Printf("Processed: %d, Failures: %d\n", totalProcessed, totalFailures)

        // Sleep for 1 second
        time.Sleep(1 * time.Second)
    }

    if err := cursor.Err(); err != nil {
        return err
    }

    log.Printf("Update completed. Total processed: %d, Total failures: %d\n", totalProcessed, totalFailures)
    return nil
}


func main() {
    // Set client options and connect to MongoDB
    clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Ensure connection is established
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")

    // Update all documents in the collection
    err = updateAllContentEmbeddings(client)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("All content embeddings updated successfully >>>")
}
