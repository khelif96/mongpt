package gpt

import (
	"context"
	"fmt"
	"os"

	"log"

	"github.com/khelif96/mongpt/userInput"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
)

// get key from env file
var client *openai.Client

const promptThreshold = 1000

func Init() {
	client = openai.NewClient(os.Getenv("OPENAI_KEY"))

}

func trainOnSchema(schemas string) ([]openai.ChatCompletionMessage, error) {
	if client == nil {
		panic("Client is nil! Did you call Init() first?")
	}
	trainingPromptTokenCount := calculateTokenUsage(initialTrainingPrompt, schemaDefinitionPrompt, schemas)
	if trainingPromptTokenCount > promptThreshold {
		errorMessage := fmt.Sprintf("Training prompt cost %d exceeds threshold %d\n It will cost approximately $%f to run.", trainingPromptTokenCount, promptThreshold, calculateGPT3Cost(trainingPromptTokenCount))
		if !userInput.PromptForAllowingExpensiveQueries(errorMessage) {
			return nil, fmt.Errorf(errorMessage)
		}
	}

	chatCompletionRequest := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: initialTrainingPrompt,
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: schemaDefinitionPrompt + schemas,
		},
	}
	return chatCompletionRequest, nil
}

func AskGPT(schemas, prompt string) (*string, error) {
	if client == nil {
		panic("Client is nil! Did you call Init() first?")
	}
	promptTokenUsage := calculateTokenUsage(prompt)
	if promptTokenUsage > promptThreshold {
		errorMessage := fmt.Sprintf("Prompt cost %d exceeds threshold %d\n It will cost approximately $%f to run.", promptTokenUsage, promptThreshold, calculateGPT3Cost(promptTokenUsage))
		if !userInput.PromptForAllowingExpensiveQueries(errorMessage) {
			return nil, fmt.Errorf(errorMessage)
		}
	}

	trainingPrompts, err := trainOnSchema(schemas)
	if err != nil {
		return nil, err
	}
	fullPrompt := append(trainingPrompts, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: fullPrompt,
		},
	)

	if err != nil {
		return nil, errors.Wrapf(err, "Error creating chat completion")
	}

	log.Println(fmt.Sprintf("Response Cost: %d, $%f", resp.Usage.TotalTokens, calculateGPT3Cost(resp.Usage.TotalTokens)))

	return &resp.Choices[0].Message.Content, nil
}