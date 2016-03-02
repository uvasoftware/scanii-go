package scaniigo

import (
	"encoding/json"
	"net/http"
)

// TempAuthTokenResponse holds the returned data from a call to the
// CreateTempAuthToken method
type TempAuthTokenResponse struct {
	ID             string `json:"id"`
	CreationDate   string `json:"creation_date"`
	ExpirationDate string `json:"expiration_date"`
}

// TempAuthTokenParams holds the parameters for calls to create a
// temp token
type TempAuthTokenParams struct {
	// Timeout is the number of seconds this token should be valid for
	// (optional defaults to 300 seconds)
	Timeout int
}

// Validate makes sure that the required parameters are present when this time is
// to be used
func (p *TempAuthTokenParams) Validate() error {
	return nil
}

// RetrieveAuthToken retrieves a previously created token
func (c *Client) RetrieveAuthToken(id string) (*TempAuthTokenResponse, error) {
	req, err := http.NewRequest("GET", c.Endpoint+AuthPath+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.APIAuth.Key, c.APIAuth.Secret)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var tatr TempAuthTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&tatr); err != nil {
		return nil, err
	}
	return &tatr, nil
}

// CreateTempAuthToken creates a temporary authentication token
func (c *Client) CreateTempAuthToken() (*TempAuthTokenResponse, error) {
	req, err := http.NewRequest("POST", c.Endpoint+AuthPath, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.APIAuth.Key, c.APIAuth.Secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var tatr TempAuthTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&tatr); err != nil {
		return nil, err
	}
	return &tatr, nil
}

// DeleteTempAuthToken deletes a temporary authentication token
func (c *Client) DeleteTempAuthToken(id string) error {
	req, err := http.NewRequest("DELETE", c.Endpoint+AuthPath+"/"+id, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.APIAuth.Key, c.APIAuth.Secret)

	_, err = c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
