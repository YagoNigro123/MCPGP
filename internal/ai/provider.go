package ai

type AIProvider interface {
	Generate(prompt string) (string, error)
}

func GetProvider(name string, apiKey string) AIProvider {
	switch name {
	case "groq":
		return &GroqProvider{ApiKey: apiKey, Model: "mixtral-8x7b-32768"}
	case "gemini":
		return &GeminiProvider{ApiKey: apiKey, Model: "gemini-pro"}
	default:
		return nil
	}
}
