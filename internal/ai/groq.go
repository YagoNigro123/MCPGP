package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	url := "https://api.groq.com/openai/v1/chat/completions"

	reqBody := groqRequest{
		Model: g.Model,
		Messages: []message{
			{Role: "user", Content: prompt},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to create JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to establish connection with Groq: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Groq error (status %d): %s", resp.StatusCode, string(body))
	}

	var result groqResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to read the response JSON: %v", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("Groq did not return any response")
	}

	return result.Choices[0].Messages.Content, nil
}
