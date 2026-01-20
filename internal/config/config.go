package config

import (
	"fmt"
	"os"
)

type Config struct {
	GroqAPIKey   string
	GeminiAPIKey string
}

func LoadConfig() (*Config, error) {
	groqKey := os.Getenv("GROQ_API_KEY")

	if groqKey == "" {
		return nil, fmt.Errorf("falta la variable de entorno GROQ_API_KEY")
	}

	geminiKey := os.Getenv("GEMINI_API_KEY")

	return &Config{
		GroqAPIKey:   groqKey,
		GeminiAPIKey: geminiKey,
	}, nil
}
