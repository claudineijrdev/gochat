package interfaces

import "context"

type IChatClient interface {
	StartChat() error
	SendMessageStream(ctx context.Context, msg string) error
	SendMessage(ctx context.Context, msg string)  (interface{}, error)
	PrintResponse(res interface{})
	Close()
}