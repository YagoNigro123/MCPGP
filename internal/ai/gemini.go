package ai

type GeminiProvider struct {
	ApiKey string
	Model  string
}

func (g *GeminiProvider) Generate(prompt string) (string, error) {

	// ...

	return "Respuesta de Groq", nil
}
