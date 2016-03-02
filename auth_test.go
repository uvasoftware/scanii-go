package scaniigo

import (
	"reflect"
	"testing"
)

// TestCreateTempAuthToken
func TestCreateTempAuthToken(t *testing.T) {
	t.Parallel()
	c, err := NewClient(&ClientOpts{
		Version:  "2.1",
		Validate: true,
	})
	if err != nil {
		t.Error(err)
	}
	tatr, err := c.CreateTempAuthToken()
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(tatr).String() != "*scaniigo.TempAuthTokenResponse" {
		t.Error(ErrInvalidDataType)
	}
}

// TestRetrieveAuthToken
func TestRetrieveAuthToken(t *testing.T) {
	t.Parallel()
	c, err := NewClient(&ClientOpts{
		Version:  "2.1",
		Validate: true,
	})
	if err != nil {
		t.Error(err)
	}
	tr, err := c.CreateTempAuthToken()
	if err != nil {
		t.Error(err)
	}
	_, err = c.RetrieveAuthToken(tr.ID)
	if err != nil {
		t.Error(err)
	}
}

// TestDeleteTempAuthToken
func TestDeleteTempAuthToken(t *testing.T) {
	t.Parallel()
	c, err := NewClient(&ClientOpts{
		Version:  "2.1",
		Validate: true,
	})
	if err != nil {
		t.Error(err)
	}
	tr, err := c.CreateTempAuthToken()
	if err != nil {
		t.Error(err)
	}
	if err := c.DeleteTempAuthToken(tr.ID); err != nil {
		t.Error(err)
	}
}
