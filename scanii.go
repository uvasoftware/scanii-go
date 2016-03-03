package scaniigo

import (
	"crypto/tls"
	"errors"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://api.scanii.com"
const basePath = "/v"

// PingPath contains the path to the ping resource
const PingPath = "/ping"

// FilePath contains the path to the files resource
const FilePath = "/files"

// FileAsyncPath contains the path to the files async resource
const FileAsyncPath = FilePath + "/async"

// FileFetchPath contains the path to the files fetch resource
const FileFetchPath = FilePath + "/fetch"

// AuthPath contains the path to the auth tokens resource
const AuthPath = "/auth/tokens"

const clientTimeout = 30

// APIAuth holds the pieces needed to authenticate against the API
type APIAuth struct {
	Key    string
	Secret string
}

// Client holds the current client settings
type Client struct {
	Endpoint   string
	APIAuth    *APIAuth
	HTTPClient *http.Client
}

// ClientOpts holds the options to build a client
type ClientOpts struct {
	Version  string
	Validate bool
}

// Validator is an interface containing a Validate method
type Validator interface {
	Validate() error
}

// RequestGenerator is an interface to be used to generate HTTP.Request types
type RequestGenerator interface {
	Generate(c *Client, execType string) (*http.Request, error)
}

// NewClient creates a new reference to a Client
func NewClient(co *ClientOpts) (*Client, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	c := &Client{
		Endpoint: baseURL + basePath + co.Version,
		HTTPClient: &http.Client{
			Timeout:   time.Duration(clientTimeout * time.Second),
			Transport: tr,
		},
	}
	auth, err := getAuth()
	if err != nil {
		return nil, err
	}
	c.APIAuth = auth
	if co.Validate {
		if _, err := c.Ping(); err != nil {
			return nil, err
		}
		return c, nil
	}
	return c, nil
}

// getAuth verifies there's enough valid
func getAuth() (*APIAuth, error) {
	key := os.Getenv("SCANII_API_KEY")
	secret := os.Getenv("SCANII_API_SECRET")
	a := &APIAuth{
		Key:    key,
		Secret: secret,
	}
	if !ValidAuth(a) {
		return nil, ErrInvalidAuth
	}
	return a, nil
}

// Validate checks to make sure the given auth is valid
func (a *APIAuth) Validate() error {
	if len(a.Key) == 32 && len(a.Secret) == 9 {
		return errors.New("")
	}
	return errors.New("")
}

// ValidAuth checks to make sure the given auth is valid
func ValidAuth(a *APIAuth) bool {
	if len(a.Key) == 32 && len(a.Secret) == 9 {
		return true
	}
	return false
}

// ConvertDate converts a string to an instance of time.Time
func ConvertDate(dt string) (time.Time, error) {
	sd, err := time.Parse(time.RFC3339Nano, dt)
	if err != nil {
		return time.Time{}, err
	}
	return sd, nil
}

// Validate runs the Validate method on the type passed in
// assuming the passed in type implements the Validator interface
func Validate(p Validator) error {
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}

// GenerateFileAPIRequest build an http.Request from the given arguments. The second argument has to
// be a type that implements the RequestGenerator interface
func GenerateFileAPIRequest(c *Client, rg RequestGenerator, execType string) (*http.Request, error) {
	return rg.Generate(c, execType)
}
