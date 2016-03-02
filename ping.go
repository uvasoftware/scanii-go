package scaniigo

import (
	"encoding/json"
	"net/http"
)

// PingResponse contains the return data from a ping request
type PingResponse struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}

// Ping makes a call to the ping endpoint
func (c *Client) Ping() (*PingResponse, error) {
	req, err := http.NewRequest("GET", c.Endpoint+PingPath, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.APIAuth.Key, c.APIAuth.Secret)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pr PingResponse
	if err := json.NewDecoder(res.Body).Decode(&pr); err != nil {
		return nil, err
	}
	return &pr, nil
}
