package scanii

import (
	"errors"
	"fmt"
	"net/http"
	"scanii-go/pkg/scanii/endpoints"
)

// Retrieves a previously processed file resource - https://docs.scanii.com/v2.1/resources.html#files
func (c *Client) Ping() (bool, error) {

	req, err := http.NewRequest("GET", endpoints.Resolve(c.Target, "ping"), nil)
	if err != nil {
		return false, err
	}
	req.Header.Set(userAgentHeader, c.UserAgentHeader)
	req.Header.Set(authorizationHeader, c.AuthenticationHeader)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, errors.New(fmt.Sprintf("HTTP error: %s", res.Status))
}
