package chat_service

import (
	"context"

	"github.com/claudineijrdev/gochat/internal/interfaces"
)

type IChatService interface{
	Start(ctx context.Context)
}

type ChatService struct {
	IChatService
	ChatClient interfaces.IChatClient
}

func NewChatService(chatClient interfaces.IChatClient) *ChatService {
	return &ChatService{ChatClient: chatClient}
}

func (cs *ChatService) Start(ctx context.Context) error {
	cs.ChatClient.StartChat()

	// response, err  := cs.ChatClient.SendMessage(ctx, "Can you name some brands of cars?")
	// if err != nil {
	// 	return  err
	// }

	// cs.ChatClient.PrintResponse(response)
	return nil
}