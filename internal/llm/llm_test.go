package llm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChatClient(t *testing.T) {
	// Setup
	// If there is any setup to be done before testing, do it here
	client, err := NewChatClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Client)
	assert.NotNil(t, client.ctx)
}

func TestResponse(t *testing.T) {
	client, err := NewChatClient()
	assert.NoError(t, err)

	res, err := client.SendWithResponse("hi, how are you")
	assert.NoError(t, err)

	assert.NotEmpty(t, res)
	t.Logf("Received response: %s\n", res)
}
