package main

import (
	"fmt"
	"github.com/abdulahmad586/alaramma.git/embeddings"
)

func main() {
    emb := embeddings.NewEmbedding("YOUR_OPENAI_API_KEY", "text-similarity-babbage-001") // Replace with your actual API key and desired model
    embeddings, err := emb.GetEmbeddings("Hello, world!")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Embeddings:", embeddings)
}
