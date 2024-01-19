package bot

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/ayush6624/go-chatgpt"

	"github.com/kaashmonee/signallm/internal/llm"
	"github.com/kaashmonee/signallm/internal/util"
)

type Bot struct {
	Number string
	Name   string
	Client *llm.ChatClient
}

type GroupInfo struct {
	GroupID string `json:"groupId"`
}

type DataMessage struct {
	Timestamp int64     `json:"timestamp"`
	Message   string    `json:"message"`
	GroupInfo GroupInfo `json:"groupInfo"`
}

type Envelope struct {
	SourceNumber string      `json:"sourceNumber"`
	DataMessage  DataMessage `json:"dataMessage"`
}

type Event struct {
	Envelope Envelope `json:"envelope"`
}

func NewBot() (*Bot, error) {
	chatClient, err := llm.NewChatClient()
	if err != nil {
		return nil, err
	}

	return &Bot{
		Number: "+12069840296",
		Name:   "mike ox",
		Client: chatClient,
	}, nil
}

func unmarshalToEvents(byteArray []byte) ([]Event, error) {
	if len(byteArray) == 0 {
		return []Event{}, nil
	}

	// There may be multiple events separated by newlines
	eventByteArrays := bytes.Split(byteArray, []byte("\n"))
	events := make([]Event, 0, len(eventByteArrays))

	for _, eventArray := range eventByteArrays {
		if len(eventArray) == 0 {
			continue
		}
		e := Event{}
		err := json.Unmarshal(eventArray, &e)
		if err != nil {
			log.Printf("Unable to unmarshal %s to Event", eventArray)
			return []Event{}, err
		}
		events = append(events, e)
	}

	return events, nil
}

func (b *Bot) receive() ([]Event, error) {
	cmd := exec.Command("java", "-jar", "/Users/kaashmonee/signal-cli/build/libs/signal-cli-fat-0.12.8-SNAPSHOT.jar", "-u", b.Number, "-o", "json", "receive")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []Event{}, nil
	}

	return unmarshalToEvents(out)
}

func (b *Bot) sendMessage(message, recipient, groupID string) error {
	cmdString := []string{"-jar", "/Users/kaashmonee/signal-cli/build/libs/signal-cli-fat-0.12.8-SNAPSHOT.jar", "-u", b.Number, "send", "-m", message}
	if groupID != "" {
		cmdString = append(cmdString, "-g", groupID)
	} else {
		cmdString = append(cmdString, recipient)
	}

	log.Printf("executing with command: %v", cmdString)
	cmd := exec.Command("java", cmdString...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	log.Println("result:", string(out))
	return nil
}

func (b *Bot) generateResponse(evnt Event) (string, error) {
	// If the message doens't contain the name of the AI or only contains the name of the AI with
	// nothing else,then ignore the message
	msg := util.RemoveName(evnt.Envelope.DataMessage.Message, b.Name)
	if strings.Trim(msg, " ") == "" {
		return "", nil
	}

	identifier := evnt.Envelope.DataMessage.GroupInfo.GroupID
	if identifier == "" {
		identifier = evnt.Envelope.SourceNumber
	}

	resp, err := b.Client.SendWithResponseAndModel(llm.SendRequest{
		Message: msg,
		Model:   chatgpt.GPT4_0613,
		User:    identifier,
	})
	if err != nil {
		return "", err
	}

	return resp, nil
}

func (b *Bot) Start() {
	for {
		time.Sleep(time.Second)
		// See if there are any messages
		events, err := b.receive()
		if err != nil {
			log.Printf("unable to receive or unmarshal response with error: %v\n", err)
			continue
		}

		for _, evnt := range events {
			response, err := b.generateResponse(evnt)
			if err != nil {
				log.Printf("Unable to generate a response for event: %v", evnt)
				continue
			}

			if response == "" {
				continue
			}

			err = b.sendMessage(response, evnt.Envelope.SourceNumber, evnt.Envelope.DataMessage.GroupInfo.GroupID)
			if err != nil {
				log.Printf("encountered error while trying to execute command with error: %v", err)
			}
		}
	}
}
