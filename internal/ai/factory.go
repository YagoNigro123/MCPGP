package ai

import (
	"fmt"

	"github.com/YagoNigro123/MCPGP/internal/config"
)

func NewAIProvider(cfg *config.Config) (AIProvider, error) {
	switch cfg.AIProvider {
	case "groq":
		return &GroqProvider{
			ApiKey: cfg.GroqAPIKey,
			Model:  cfg.GroqModel,
		}, nil
	case "gemini":
		return &GeminiProvider{
			ApiKey: cfg.GeminiAPIKey,
			Model:  cfg.GeminiModel,
		}, nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", cfg.AIProvider)
	}
}
