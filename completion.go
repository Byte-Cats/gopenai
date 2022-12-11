package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const openaiURL = "https://api.openai.com/v1/completions"

type CompletionsRequest struct {
	Model         string  `json:"model"`
	Prompt        string  `json:"prompt"`
	Temperature   float64 `json:"temperature"`
	MaxTokens     int     `json:"max_tokens"`
	TopP          float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty float64 `json:"presence_penalty"`
}

type CompletionsResponse struct {
	ID        string `json:"id"`
	Model     string `json:"model"`
	Choices   []struct {
		Text  string `json:"text"`
		Logprobs []float64 `json:"logprobs"`
		Finished bool `json:"finished"`
		Score float64 `json:"score"`
	} `json:"choices"`
}

func getCompletions(req *CompletionsRequest) (*CompletionsResponse, error) {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	var resp CompletionsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func main() {
	req := &CompletionsRequest{
		Model: "code-davinci-002",
		Prompt: "// create a header that is gradient purple with navbar home and contact us",
		Temperature: 0.22,
		MaxTokens: 256,
		TopP: 1,
		FrequencyPenalty: 0,
		PresencePenalty: 0,
	}

	resp, err := getCompletions(req)
	if err != nil {
		log.Fatal
