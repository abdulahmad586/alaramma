package embeddings

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// embedding struct to hold API key and model
type embedding struct {
    apiKey string
    model  string
}

// requestBody struct to format the JSON request body
type requestBody struct {
    Model   string   `json:"model"`
    Inputs  []string `json:"inputs"`
}

// embeddingResponse struct to parse the JSON response
type embeddingResponse struct {
    Data []struct {
        Embedding []float64 `json:"embedding"`
    } `json:"data"`
}

// NewEmbedding creates a new instance of embedding. 
// This function requires two arguments: 
// - apiKey: A string representing the OpenAI API key. This is necessary for authentication when making API requests.
// - model: A string indicating the specific OpenAI model to be used for generating embeddings. 
//   Examples include "text-similarity-babbage-001", "text-similarity-ada-001", etc.
// The function returns a pointer to an embedding instance which can be used to call the GetEmbeddings method.
//
// Example:
//   emb := embeddings.NewEmbedding("your-api-key", "text-similarity-babbage-001")
//   embeddings, err := emb.GetEmbeddings("Your text here")
func NewEmbedding(apiKey, model string) *embedding {
    return &embedding{apiKey: apiKey, model: model}
}

// getEmbeddings method to convert text to embeddings
func (e *embedding) GetEmbeddings(text string) ([]float64, error) {
    url := "https://api.openai.com/v1/embeddings"

    reqBody := requestBody{
        Model:  e.model, // use the model specified in the embedding instance
        Inputs: []string{text},
    }

    jsonReq, err := json.Marshal(reqBody)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", "Bearer "+e.apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    var respEmbedding embeddingResponse
    if err := json.Unmarshal(body, &respEmbedding); err != nil {
        return nil, err
    }

    if len(respEmbedding.Data) > 0 {
        return respEmbedding.Data[0].Embedding, nil
    }

    return nil, fmt.Errorf("no embeddings returned");
}
