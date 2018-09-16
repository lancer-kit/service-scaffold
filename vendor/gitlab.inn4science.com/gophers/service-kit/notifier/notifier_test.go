package notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSender(t *testing.T) {
	c, e := NewSender(nil)
	if !assert.NoError(t, e) {
		return
	}
	if !assert.NotEmpty(t, c) {
		return
	}
	e = c.IsConnected()
	if !assert.NoError(t, e) {
		return
	}
	e = c.Send(&Message{
		UserId:      0,
		Command:     "test",
		IsBroadcast: true,
	})
	assert.NoError(t, e)
	c.Disconnect()
}
