package scanii

import (
	"encoding/json"
	"errors"
	"github.com/uvasoftware/scanii-go/endpoints"
	"github.com/uvasoftware/scanii-go/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// RetrieveAuthToken retrieves a previously created token
func (c *Client) RetrieveAuthToken(id string) (*models.AuthToken, error) {

	req, err := http.NewRequest("GET", endpoints.Resolve(c.Target, "auth/tokens/")+id, nil)
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

	content, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(content))
	}

	var r models.AuthToken
	if err := json.Unmarshal(content, &r); err != nil {
		return nil, err
	}
	return &r, nil

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
	req.Header.Set(contentTypeHeader, "application/x-www-form-urlencoded")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusCreated {
		return nil, errors.New(string(content))
	}

	var r models.AuthToken
	if err := json.Unmarshal(content, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// DeleteTempAuthToken deletes a temporary authentication token
func (c *Client) DeleteAuthToken(id string) error {
	req, err := http.NewRequest("DELETE", endpoints.Resolve(c.Target, "auth/tokens/")+id, nil)
	if err != nil {
		return err
	}

	req.Header.Set(userAgentHeader, c.UserAgentHeader)
	req.Header.Set(authorizationHeader, c.AuthenticationHeader)

	res, err := c.HTTPClient.Do(req)

	content, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusNoContent {
		return errors.New(string(content))
	}

	if err != nil {
		return err
	}
	return nil
}
