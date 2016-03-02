package scaniigo

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://api.scanii.com"
const basePath = "/v"

const (
	PingPath = "/ping"

	FilePath      = "/files"
	FileAsyncPath = FilePath + "/async"
	FileFetchPath = FilePath + "/fetch"

	AuthPath = "/auth/tokens"
)

const clientTimeout = 30

// Layout holds the default layout for timestamps
const Layout = "2006-01-02T15:04"

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

// ValidAuth checks to make sure the given auth is valid
func ValidAuth(a *APIAuth) bool {
	if len(a.Key) == 32 && len(a.Secret) == 9 {
		return true
	}
	return false
}

// ConvertDate converts a string to an instance of time.Time
func ConvertDate(dt string) (time.Time, error) {
	sd, err := time.Parse(Layout, dt)
	if err != nil {
		return time.Time{}, err
	}
	return sd, nil
}
