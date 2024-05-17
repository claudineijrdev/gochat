package openai_client

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/claudineijrdev/gochat/internal/interfaces"
	"github.com/sashabaranov/go-openai"
)

type OpenAiClient struct {
	ModelName string
	interfaces.IChatClient
	client *openai.Client
	ctx context.Context
	request *openai.ChatCompletionRequest
}

func NewOpenAiClient(ctx context.Context, token, model string) (*OpenAiClient, error) {
	if token == "" {
		return nil, fmt.Errorf("token cannot be empty")
	}

	if model == "" {
		return nil, fmt.Errorf("model cannot be empty")
	}

	client := openai.NewClient(token)

	return &OpenAiClient{
		ModelName: model,
		client: client,
		ctx: ctx,
	}, nil
}

func (ga *OpenAiClient) StartChat() error {
	if ga.client == nil {
		return  fmt.Errorf("client not initialized")
	}

	ga.request = &openai.ChatCompletionRequest{
		Model: ga.ModelName,
	}
	return nil
	
}

func (ga *OpenAiClient) SendMessageStream(ctx context.Context, msg string) error {
	ga.request.MaxTokens = 20
	ga.request.Messages = append(ga.request.Messages, openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleUser,
		Content: msg,
	})
	ga.request.Stream = true


	stream, err := ga.client.CreateChatCompletionStream(ctx, *ga.request)

	if err != nil {
		return err
	}
	defer stream.Close()

	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF){
			break
		}
		if err != nil {
			return err
		}

		fmt.Println(res.Choices[0].Delta.Content)
	}

	return nil
}

func (ga *OpenAiClient) SendMessage(ctx context.Context, msg string)  (interface{}, error) {
	ga.request.Messages = append(ga.request.Messages, openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleUser,
		Content: msg,
	})


	resp, err := ga.client.CreateChatCompletion(ctx, *ga.request)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ga *OpenAiClient) PrintResponse(res interface{}) {
	fmt.Println(res.(*openai.ChatCompletionResponse).Choices[0].Message.Content)
}

func (ga *OpenAiClient) Close() {
	
}
