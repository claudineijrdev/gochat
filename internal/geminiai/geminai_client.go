package geminiai

import (
	"context"
	"fmt"

	"cloud.google.com/go/vertexai/genai"
	"github.com/claudineijrdev/gochat/internal/interfaces"
	"google.golang.org/api/iterator"
)

type GeminiAiClient struct {
	ModelName string
	interfaces.IChatClient
	client *genai.Client
	model *genai.GenerativeModel
	session *genai.ChatSession
}

func NewGeminiAiClient(ctx context.Context, modelName, projectID, location string) (*GeminiAiClient, error) {
	client, err := genai.NewClient(ctx,  projectID, location)
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel(modelName)

	return &GeminiAiClient{
		ModelName: modelName,
		client: client,
		model: model,
	}, nil
}

func (ga *GeminiAiClient) StartChat() error {
	if ga.client == nil {
		return  fmt.Errorf("client not initialized")
	}

	ga.session = ga.model.StartChat()
	return nil
	
}

func (ga *GeminiAiClient) SendMessageStream(ctx context.Context, msg string) error {
	iter := ga.session.SendMessageStream(ctx, genai.Text(msg))
	for {
		res, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		ga.PrintResponse(res)
	}
	return nil
}

func (ga *GeminiAiClient) SendMessage(ctx context.Context, msg string)  (interface{}, error) {
	res,err := ga.session.SendMessage(ctx, genai.Text(msg))
	if err != nil {
		return nil, err
	}	
	return res, nil
}

func (ga *GeminiAiClient) PrintResponse(res interface{}) {
	for _, cand := range res.(*genai.GenerateContentResponse).Candidates {
		for _, part := range cand.Content.Parts {
			fmt.Println(part)
		}
	}
}

func (ga *GeminiAiClient) Close() {
	ga.client.Close()
}
