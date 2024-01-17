package bot

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/kaashmonee/signallm/internal/llm"
)

type Bot struct {
	Number       string
	msgsReceived chan string
	Client       *llm.ChatClient
}

func NewBot() *Bot {
	chatClient, err := llm.NewChatClient()
	if err != nil {
		log.Fatalf("Unable to initialize bot with error: %s", err.Error())
		return nil
	}

	return &Bot{
		Number:       "+12069840296",
		msgsReceived: make(chan string),
		Client:       chatClient,
	}
}

func receive(number string) string {
	cmd := exec.Command("java", "-jar", "/Users/kaashmonee/signal-cli/build/libs/signal-cli-fat-0.12.8-SNAPSHOT.jar", "-u", "+12069840296", "receive")
	out, err := cmd.CombinedOutput()
	if err != nil {
		// log.Fatalf("cmd.Run() failed with %s\n", err)
		return err.Error()
	}
	// fmt.Printf("Combined output:\n%s\n", string(out))
	return string(out)
}

func sendMessage(message string) error {
	cmd := exec.Command("java", "-jar", "/Users/kaashmonee/signal-cli/build/libs/signal-cli-fat-0.12.8-SNAPSHOT.jar", "-u", "+12069840296", "send", "-m", message, "+12246000039")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println("result:", string(out))
	return nil
}

func (b *Bot) generateResponse(message string) (string, error) {
	resp, err := b.Client.SendWithResponse(message)
	if err != nil {
		return "", err
	}

	return resp, nil
}

func (b *Bot) Start() {
	for {
		time.Sleep(time.Second)
		// See if there are any messages
		out := receive(b.Number)
		fmt.Printf("received the message: %s\n", out)
		response, err := b.generateResponse(out)
		if err != nil {
			response = fmt.Sprintf("Unable to obtain response with error: %s", err.Error())
		}

		go sendMessage(response)
	}
}
