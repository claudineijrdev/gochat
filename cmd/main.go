package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/claudineijrdev/gochat/infra/config"
	"github.com/claudineijrdev/gochat/internal/factories"
	chat_service "github.com/claudineijrdev/gochat/internal/service"
)


func init(){
	config.LoadConfig()
}

func main(){
	ctx := context.Background()

	chatClient,err := factories.NewClientFactory().Create("openai", ctx)
	if err != nil {
		log.Fatalf("Failed to create ChatClient: %v", err)
	}
		
	service := chat_service.NewChatService(chatClient)

	err = terminal(service)

	if err != nil {
		log.Fatalf("Failed to run terminal: %v", err)
	}
}

func terminal(chat *chat_service.ChatService) error {
	chat.Start(context.Background())
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to GoChat! (Type 'exit' to quit)")
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %v", err)
		}
		filteredInput := strings.TrimSpace(input)
		if filteredInput == "exit" {
			fmt.Println("Exiting...")
			break
		}
		//err = chat.ChatClient.SendMessageStream(context.Background(), filteredInput)
		resp, err := chat.ChatClient.SendMessage(context.Background(), filteredInput)
		if err != nil {
			return fmt.Errorf("failed to send message: %v", err)
		}
		chat.ChatClient.PrintResponse(resp)

	}
	return nil
}