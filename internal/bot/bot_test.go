package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBotInit(t *testing.T) {
	t.Log("Testing bot initialization")
	bot, err := NewBot()
	if err != nil {
		t.Errorf("error initializing bot: %v", err)
	}

	assert.NotEmpty(t, bot.Number)
}

func TestUnmarshalToEvent(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected Event
	}{
		{"empty", []byte{}, Event{}},
		{"contains message", []byte(`{"envelope":{"source":"+12246000039","sourceNumber":"+12246000039","sourceUuid":"dc734153-986d-4278-a8be-e82e9b62cd78","sourceName":"Skanda","sourceDevice":1,"timestamp":1705594373870,"dataMessage":{"timestamp":1705594373870,"message":"hi","expiresInSeconds":0,"viewOnce":false,"groupInfo":{"groupId":"lx5vjquCWGByxrCFSDHf4C1W1XOjdH1m+JGz8bIJIBc=","type":"DELIVER"}}},"account":"+12069840296"}`), Event{Envelope: Envelope{DataMessage: DataMessage{Message: "hi"}}}},
		{"does not contain message", []byte(`{"envelope":{"source":"+12246000039","sourceNumber":"+12246000039","sourceUuid":"dc734153-986d-4278-a8be-e82e9b62cd78","sourceName":"Skanda","sourceDevice":1,"timestamp":1705594373616,"typingMessage":{"action":"STARTED","timestamp":1705594373616,"groupId":"lx5vjquCWGByxrCFSDHf4C1W1XOjdH1m+JGz8bIJIBc="}},"account":"+12069840296"}`), Event{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := unmarshalToEvents(tt.input)
			assert.NoError(t, err)
			for _, result := range results {
				switch tt.name {
				case "empty", "does not contain message":
					assert.Equal(t, result.Envelope.DataMessage.Message, "")
				case "contains message":
					assert.Equal(t, result.Envelope.DataMessage.Message, "hi")
				}
			}
		})
	}
}

func TestSendToSingleRecipient(t *testing.T) {
	bot, err := NewBot()
	assert.Nil(t, err)

	err = bot.sendMessage("hi", "+12246000039", "")
	assert.Nil(t, err, "encountered error: %v", err)
}

func TestSendToGroup(t *testing.T) {
	bot, err := NewBot()
	assert.Nil(t, err)

	err = bot.sendMessage("hi", "", "lx5vjquCWGByxrCFSDHf4C1W1XOjdH1m+JGz8bIJIBc=")
	assert.Nil(t, err)
}
