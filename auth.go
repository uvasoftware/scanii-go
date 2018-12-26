package scanii

import (
	"encoding/json"
	"github.com/uvasoftware/scanii-go/endpoints"
	"github.com/uvasoftware/scanii-go/models"
	"net/http"
	"net/url"
	"strings"
)

// RetrieveAuthToken retrieves a previously created token
func (c *Client) RetrieveAuthToken(id string) (*models.AuthToken, error) {

	req, err := http.NewRequest("DELETE", endpoints.Resolve(c.Target, "auth/tokens/")+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(userAgentHeader, c.UserAgentHeader)
	req.Header.Set(authorizationHeader, c.AuthenticationHeader)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response models.AuthToken
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil

}

// CreateTempAuthToken creates a temporary authentication token
func (c *Client) CreateAuthToken(timeout int) (*models.AuthToken, error) {

	if timeout <= 0 {
		timeout = 5
	}

	options := url.Values{}
	options.Set("timeout", string(timeout))
	req, err := http.NewRequest("POST", endpoints.Resolve(c.Target, "auth/tokens"), strings.NewReader(options.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set(userAgentHeader, c.UserAgentHeader)
	req.Header.Set(authorizationHeader, c.AuthenticationHeader)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response models.AuthToken
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteTempAuthToken deletes a temporary authentication token
func (c *Client) DeleteAuthToken(id string) error {
	req, err := http.NewRequest("DELETE", endpoints.Resolve(c.Target, "auth/tokens/")+id, nil)
	if err != nil {
		return err
	}

	req.Header.Set(userAgentHeader, c.UserAgentHeader)
	req.Header.Set(authorizationHeader, c.AuthenticationHeader)

	_, err = c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
