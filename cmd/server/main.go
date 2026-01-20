package main

import (
	"fmt"
	"log"

	"github.com/YagoNigro123/MCPGP/internal/config"
)

func main() {
	fmt.Println("Start Server MCP...")

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error config: %v", err)
	}

	fmt.Println("Config charge sucessfuly.")
	fmt.Printf("Groq Key detected (long: %d characters)\n", len(cfg.GroqAPIKey))

	// ...

}
