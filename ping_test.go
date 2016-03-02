package scaniigo

import (
	"reflect"
	"testing"
)

// TestPing makes sure *Client.Ping() behaves as it should
func TestPing(t *testing.T) {
	t.Parallel()
	clientOpts := &ClientOpts{
		Version:  "2.1",
		Validate: true,
	}
	c, err := NewClient(clientOpts)
	if err != nil {
		t.Error(err)
	}
	res, err := c.Ping()
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(res).String() != "*scaniigo.PingResponse" {
		t.Error(ErrInvalidDataType)
	}
}
