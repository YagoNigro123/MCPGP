package main

import (
	"log"

	"github.com/YagoNigro123/MCPGP/internal/ai"
	"github.com/YagoNigro123/MCPGP/internal/config"
)

func main() {
	log.Println("Starting MCP server...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	groq := &ai.GroqProvider{
		ApiKey: cfg.GroqAPIKey,
		Model:  "llama-3.3-70b-versatile",
	}

	prompt := "Why is Go language so fast?"
	log.Printf("Question to Groq: '%s'...", prompt)

	res, err := groq.Generate(prompt)
	if err != nil {
		log.Fatalf("failed to get response: %v", err)
	}

	log.Println("\nResponse:")
	log.Println("--------------------------------------------------")
	log.Println(res)
	log.Println("--------------------------------------------------")
}
