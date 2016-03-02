package scaniigo

import (
	"reflect"
	"testing"
)

// TestNewClient validates a new Scanii client is created successfully
func TestNewClient(t *testing.T) {
	t.Parallel()
	c, err := NewClient(&ClientOpts{
		Version:  "2.1",
		Validate: true,
	})
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(c).String() != "*scaniigo.Client" {
		t.Error(ErrInvalidDataType)
	}
}

// TestgetAuth validates that authentication parameters can be retrieved
// successfully
func TestgetAuth(t *testing.T) {
	t.Parallel()
	a, err := getAuth()
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(a).String() != "*scaniigo.APIAuth" {
		t.Error(ErrInvalidDataType)
	}
}

// TestValidAuth validates that authentication is valid
func TestValidAuth(t *testing.T) {
	t.Parallel()
	a, err := getAuth()
	if err != nil {
		t.Error(err)
	}
	if !ValidAuth(a) {
		t.Error(ErrInvalidAuth)
	}
}

// TestConvertDate validates that date strings are successfully converted
// to the desired time.Time format
func TestConvertDate(t *testing.T) {
	t.Parallel()
	_, err := ConvertDate("2015-05-29T07:18:58.123Z")
	if err != nil {
		t.Error(err)
	}
}
