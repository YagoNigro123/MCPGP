package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GroqProvider struct {
	ApiKey string
	Model  string
}

type groqRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type groqResponse struct {
	Choices []struct {
		Messages message `json:"message"`
	}

	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func (g *GroqProvider) Generate(prompt string) (string, error) {
	log.Println("Groq: generating completion")

	url := "https://api.groq.com/openai/v1/chat/completions"

	reqBody := groqRequest{
		Model: g.Model,
		Messages: []message{
			{Role: "user", Content: prompt},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Printf("Groq: failed to marshal request JSON: %v", err)
		return "", fmt.Errorf("failed to create JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Groq: failed to create HTTP request: %v", err)
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.ApiKey)

	log.Printf("Groq: sending request (model=%s)", g.Model)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Groq: request failed: %v", err)
		return "", fmt.Errorf("failed to establish connection with Groq: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("Groq: non-200 response (status=%d)", resp.StatusCode)
		return "", fmt.Errorf("Groq error (status %d): %s", resp.StatusCode, string(body))
	}

	var result groqResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Groq: failed to unmarshal response JSON: %v", err)
		return "", fmt.Errorf("failed to read the response JSON: %v", err)
	}

	if len(result.Choices) == 0 {
		log.Println("Groq: empty choices in response")
		return "", fmt.Errorf("Groq did not return any response")
	}

	log.Println("Groq: response generated successfully")
	return result.Choices[0].Messages.Content, nil
}
