package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AIProvider   string
	GroqAPIKey   string
	GroqModel    string
	GeminiAPIKey string
	GeminiModel  string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	provider := os.Getenv("AI_PROVIDER")
	if provider == "" {
		return nil, fmt.Errorf("AI_PROVIDER is required in .env")
	}

	cfg := &Config{
		AIProvider:   provider,
		GroqAPIKey:   os.Getenv("GROQ_API_KEY"),
		GroqModel:    os.Getenv("GROQ_MODEL"),
		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
		GeminiModel:  os.Getenv("GEMINI_MODEL"),
	}

	if cfg.AIProvider == "groq" && cfg.GroqAPIKey == "" {
		return nil, fmt.Errorf("GROQ_API_KEY is required when using Groq")
	}
	if cfg.AIProvider == "gemini" && cfg.GeminiAPIKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is required when using Gemini")
	}

	return cfg, nil
}
