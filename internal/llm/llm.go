package llm

import (
	"context"
	"log"
	"os"

	"github.com/ayush6624/go-chatgpt"
)

type ChatClient struct {
	*chatgpt.Client
	ctx context.Context
}

type SendRequest struct {
	Message string
	Model   chatgpt.ChatGPTModel
	User    string
}

type Message struct {
	Content string `json:"content"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Response struct {
	Choices []Choice
}

func NewChatClient() (*ChatClient, error) {
	client, err := chatgpt.NewClient(os.Getenv("OPENAI_KEY"))
	if err != nil {
		return nil, err
	}
	return &ChatClient{
		Client: client,
		ctx:    context.Background(),
	}, nil
}

func (c *ChatClient) SendWithResponse(message string) (string, error) {
	res, err := c.SimpleSend(c.ctx, message)
	if err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", nil
	}

	return res.Choices[0].Message.Content, nil
}

func (c *ChatClient) SendWithResponseAndModel(request SendRequest) (string, error) {
	req := &chatgpt.ChatCompletionRequest{
		Model: request.Model,
		Messages: []chatgpt.ChatMessage{
			{
				Role:    chatgpt.ChatGPTModelRoleUser,
				Content: request.Message,
			},
		},
		User: request.User,
	}

	res, err := c.Send(c.ctx, req)
	if err != nil {
		log.Printf("error sending message: %v", err)
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", nil
	}

	return res.Choices[0].Message.Content, nil
}
