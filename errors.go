package scaniigo

import (
	"errors"
)

// ErrInvalidKey is given when a key is invalid
var ErrInvalidKey = errors.New("invalid api key")

// ErrInvalidSecret is given when a secret is invalid
var ErrInvalidSecret = errors.New("invalid api secret")

var ErrInvalidAuth = errors.New("invalid auth parameters")

// ErrUnauthorizedRequest is given when a request receives a 401
type ErrUnauthroized struct {
	Error string
}
