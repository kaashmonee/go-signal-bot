package llm

import (
	"context"
	"encoding/json"
	"os"

	"github.com/ayush6624/go-chatgpt"
)

type ChatClient struct {
	*chatgpt.Client
	ctx context.Context
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

	formatted, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}

	return string(formatted), nil
}
