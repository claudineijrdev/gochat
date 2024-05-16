package mocks

import (
	"context"

	"github.com/claudineijrdev/gochat/internal/geminiai"
	"github.com/stretchr/testify/mock"
)

type MockGeminiAiClient struct {
	mock.Mock
	geminiai.GeminiAiClient
}

func NewMockGeminiAiClient() *MockGeminiAiClient {
	client := &MockGeminiAiClient{}
	client.On("Close").Return()
	client.On("SendMessageStream", mock.Anything, mock.Anything).Return(nil)
	client.On("SendMessage", mock.Anything, mock.Anything).Return(nil, nil)
	client.On("PrintResponse", mock.Anything).Return()
	client.On("StartChat").Return(nil)
	return client
}

func (m *MockGeminiAiClient) Close() {
	m.Called()
}

func (m *MockGeminiAiClient) SendMessageStream(ctx context.Context, msg string) error {
	args := m.Called(ctx, msg)
	return args.Error(0)
}

func (m *MockGeminiAiClient) SendMessage(ctx context.Context, msg string) (interface{}, error) {
	args := m.Called(ctx, msg)
	return args.Get(0), args.Error(1)
} 

func (m *MockGeminiAiClient)  PrintResponse(res interface{}) {
	m.Called(res)
}

func (m *MockGeminiAiClient)  StartChat() error {
	args := m.Called()
	return args.Error(0)
}






