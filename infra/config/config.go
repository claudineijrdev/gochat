package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvConf struct {
	GeminiApiKey 		  string
	GcpProjectId 		  string
	GcpLocation 		  string
	GeminiModelName 	  string

	OpenAiApiKey 		  string
	OpenAiModelName 	  string
}

var envCfg *EnvConf

func LoadConfig() *EnvConf {
	if envCfg == nil {
		envCfg = load()
	}
	return envCfg
}

func load() *EnvConf {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	cfg := new(EnvConf)
	cfg.GeminiApiKey = os.Getenv("GEMINI_API_KEY")
	cfg.GcpProjectId = os.Getenv("GCP_PROJECT_ID")
	cfg.GcpLocation = os.Getenv("GCP_LOCATION")
	cfg.GeminiModelName = os.Getenv("GEMINI_MODEL_NAME")
	cfg.OpenAiApiKey = os.Getenv("OPEN_AI_API_KEY")
	cfg.OpenAiModelName = os.Getenv("OPEN_AI_MODEL_NAME")

	return cfg
}
