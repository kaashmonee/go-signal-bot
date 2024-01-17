package bot

import (
	"fmt"
	"os/exec"
	"time"
)

type Bot struct {
	Number       string
	msgsReceived chan string
}

func NewBot() *Bot {
	return &Bot{
		Number:       "+12069840296",
		msgsReceived: make(chan string),
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
	cmd := exec.Command("java", "-jar", "/Users/kaashmonee/signal-cli/build/libs/signal-cli-fat-0.12.8-SNAPSHOT.jar", "-u", "+12069840296", "send", "-m", "\"just received a non-empty message\"", "+12246000039")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println("result:", string(out))
	return nil
}

func (b *Bot) Start() {
	for {
		time.Sleep(time.Second)
		// See if there are any messages
		out := receive(b.Number)
		fmt.Printf("received the message: %s\n", out)
		if out != "" {
			go sendMessage("wow, here's a response to a non-empty message!")
		}
	}
}
