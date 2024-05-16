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
	return nil
}