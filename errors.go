package scaniigo

import (
	"errors"
)

// ErrInvalidKey is given when a key is invalid
var ErrInvalidKey = errors.New("invalid api key")

// ErrInvalidSecret is given when a secret is invalid
var ErrInvalidSecret = errors.New("invalid api secret")

// ErrInvalidAuth is given when the authentication parameters
// are not correct
var ErrInvalidAuth = errors.New("invalid auth parameters")

// ErrInvalidDataType is given when a unit test looks to see the type
// returned from a call and it's not the one it's expecting
var ErrInvalidDataType = errors.New("invalid data type returned")

// ErrUnauthorized is given when a request receives a 401
type ErrUnauthroized struct {
	Error string
}
