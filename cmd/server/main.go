package main

import (
	"log"

	"github.com/YagoNigro123/MCPGP/internal/config"
)

func main() {
	log.Println("Starting MCP server...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	log.Println("Configuration loaded successfully")
	log.Printf("Groq API key detected (length: %d characters)", len(cfg.GroqAPIKey))

	// ...
}
