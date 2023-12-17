package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/abdulahmad586/alaramma.git/db"
	"github.com/abdulahmad586/alaramma.git/embeddings"
)

func main() {
	db.InitDB("./quran.db")
	defer db.CloseDB()

	translations, err := FetchAllTranslations()
	if err != nil {
		log.Fatal("Error fetching translations:", err)
	}

	err = GenerateEmbeddingsForTranslations(translations)
	if err != nil {
		log.Fatal("Error generating embeddings for translations:", err)
	}

	fmt.Println("Embeddings generation and update for translations completed.")
}

type Translation struct {
	ID            int
	VerseNumber   int
	ChapterNumber int
	Content       string
}

func FetchAllTranslations() ([]Translation, error) {
	rows, err := db.DB.Query("SELECT id, verseNumber, chapterNumber, content FROM translations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var translations []Translation
	for rows.Next() {
		var t Translation
		if err := rows.Scan(&t.ID, &t.VerseNumber, &t.ChapterNumber, &t.Content); err != nil {
			return nil, err
		}
		translations = append(translations, t)
	}
	return translations, nil
}
func GenerateEmbeddingsForTranslations(translations []Translation) error {
	emb := embeddings.NewEmbedding(os.Getenv("OPENAI_SECRET"), "text-embedding-ada-002")

	for _, translation := range translations {
		embeddings, err := emb.GetEmbeddings(translation.Content)
		if err != nil {
			log.Printf("Error generating embeddings for translation ID %d: %v", translation.ID, err)
			continue
		}

		err = UpdateTranslationEmbedding(translation.ID, embeddings)
		if err != nil {
			log.Printf("Error updating embeddings for translation ID %d: %v", translation.ID, err)
		}
	}
	return nil
}
func UpdateTranslationEmbedding(translationID int, embeddings []float64) error {
	embeddingJSON, err := json.Marshal(embeddings)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("UPDATE translations SET contentEmbedding = ? WHERE id = ?", embeddingJSON, translationID)
	return err
}
