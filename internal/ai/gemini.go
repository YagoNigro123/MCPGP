package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeminiProvider struct {
	ApiKey string
	Model  string
}

type geminiRequest struct {
	Contents []content `json:"contents"`
}

type content struct {
	Parts []part `json:"parts"`
}

type part struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Candidates []struct {
		Content content `json:"content"`
	} `json:"candidates"`

	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func (g *GeminiProvider) Generate(prompt string) (string, error) {

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", g.Model, g.ApiKey)
	reqBody := newGeminiBody(prompt)

	respBytes, err := g.postJSON(url, reqBody)
	if err != nil {
		return "", err
	}

	return parseGeminiResponse(respBytes)
}

func newGeminiBody(prompt string) geminiRequest {
	return geminiRequest{
		Contents: []content{{Parts: []part{{Text: prompt}}}},
	}
}

func (g *GeminiProvider) postJSON(url string, data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("JSON marshal error: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error (status %d): %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func parseGeminiResponse(data []byte) (string, error) {
	var result geminiResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("JSON unmarshal error: %w", err)
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response from Gemini")
	}

	return result.Candidates[0].Content.Parts[0].Text, nil
}
