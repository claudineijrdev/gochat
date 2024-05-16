package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	geminiai "github.com/claudineijrdev/gochat/internal/geminiai"
	chat_service "github.com/claudineijrdev/gochat/internal/service"
	"github.com/joho/godotenv"
)

var (
	modelName string
	location string
	projectID string
)
func init(){
	godotenv.Load()
	modelName = os.Getenv("MODEL_NAME")
	location = os.Getenv("GCP_LOCATION")
	projectID = os.Getenv("GCP_PROJECT_ID")
}

func main(){
	ctx := context.Background()

	geminiAiClient,err := geminiai.NewGeminiAiClient(ctx, modelName, projectID, location)
	if err != nil {
		log.Fatalf("Failed to create GeminiAiClient: %v", err)
	}
	defer geminiAiClient.Close()

	service := chat_service.NewChatService(geminiAiClient)

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
		err = chat.ChatClient.SendMessageStream(context.Background(), filteredInput)
		if err != nil {
			return fmt.Errorf("failed to send message: %v", err)
		}

	}
	return nil
}