package scanii

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

// ErrFileFieldEmpty is given when a param type has an empty file field
var ErrFileFieldEmpty = errors.New("file field empty")

// ErrLocationFieldEmpty is given when a param type has an empty location field
var ErrLocationFieldEmpty = errors.New("location field empty")

// ErrUnrecognizedExecType is given when execType is used and doesn't match
// "sync" or "async"
var ErrUnrecognizedExecType = errors.New("unrecognized execType")

// ErrUnauthorized is given when a request receives a 401
type ErrUnauthorized struct {
	Error string
}
