package factories

import (
	"context"
	"fmt"

	"github.com/claudineijrdev/gochat/infra/config"
	geminiai_client "github.com/claudineijrdev/gochat/internal/geminiai"
	"github.com/claudineijrdev/gochat/internal/interfaces"
	openai_client "github.com/claudineijrdev/gochat/internal/openai"
	"github.com/sashabaranov/go-openai"
)

type Factory struct{
}
func NewClientFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Create(clientType string, ctx context.Context) (interfaces.IChatClient, error) {
	switch clientType {
	case "gemini":
		return f.createGemini(ctx)
	case "openai":
		return f.createOpenAi(ctx)
	default:
		return nil, fmt.Errorf("client type not supported")
	}
}

func (g *Factory) createGemini(ctx context.Context) (*geminiai_client.GeminiAiClient, error) {
	env := config.LoadConfig()
	modelName := env.GeminiModelName
	location := env.GcpLocation
	projectID := env.GcpProjectId

	client, err :=  geminiai_client.NewGeminiAiClient( ctx, modelName, projectID, location)
	if err != nil {
		return nil, fmt.Errorf("failed to create GeminiAiClient: %v", err.Error())
	}
	return client, nil
}

func (f *Factory) createOpenAi(ctx context.Context) (*openai_client.OpenAiClient, error) {
	env := config.LoadConfig()
	token := env.OpenAiApiKey
	model := openai.GPT3Dot5Turbo  

	client, err := openai_client.NewOpenAiClient(ctx, token, model)
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenAiClient: %v", err.Error())
	}
	return client, nil
}

