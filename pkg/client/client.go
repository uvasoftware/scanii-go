package client

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

//TODO(raf) always update this prior to a release:
var version = "2.0.0"

const userAgentHeader = "User-Agent"
const authorizationHeader = "Authorization"
const contentTypeHeader = "Content-Type"

// Client holds the current client settings
type Client struct {
	Target               string
	AuthenticationHeader string
	UserAgentHeader      string
	HTTPClient           *http.Client
}

// ClientOpts holds the options to build a client
type Opts struct {
	Target     string
	Key        string
	Secret     string
	HTTPClient *http.Client
}

// Creates a new client
func NewClient(co *Opts) *Client {
	client := new(Client)

	if co.HTTPClient == nil {
		client.HTTPClient = http.DefaultClient
	}

	client.UserAgentHeader = fmt.Sprintf("scanii-go/v%s", version)
	client.AuthenticationHeader = "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", co.Key, co.Secret)))
	client.Target = co.Target

	return client
}
