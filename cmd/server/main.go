package main

import (
	"fmt"
	"log"

	"github.com/YagoNigro123/MCPGP/internal/ai"
	"github.com/YagoNigro123/MCPGP/internal/config"
)

func main() {
	fmt.Println("Starting MCP Server...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	fmt.Printf("Configuration loaded. Selected Provider: %s\n", cfg.AIProvider)

	provider, err := ai.NewAIProvider(cfg)
	if err != nil {
		log.Fatalf("Factory error: %v", err)
	}

	prompt := "Explain in one sentence why using Environment Variables is best practice."

	fmt.Println("Sending request to AI...")
	response, err := provider.Generate(prompt)
	if err != nil {
		log.Fatalf("Generation error: %v", err)
	}

	fmt.Printf("\nResponse (%s):\n%s\n", cfg.AIProvider, response)
}
