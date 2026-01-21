package config

import (
	"errors"
	"os"
)

type Config struct {
	GroqAPIKey   string
	GeminiAPIKey string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		GroqAPIKey:   os.Getenv("GROQ_API_KEY"),
		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
	}

	if cfg.GroqAPIKey == "" {
		return nil, errors.New("GROQ_API_KEY environmet variable is required")
	}

	return cfg, nil
}
