package ai

type GroqProvider struct {
	ApiKey string
	Model  string
}

func (g *GroqProvider) Generate(prompt string) (string, error) {

	// ...

	return "Respuesta de Groq", nil
}
