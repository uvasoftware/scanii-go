package scanii

import (
	"encoding/json"
	"net/http"
	"scanii-go/pkg/scanii/endpoints"
	"scanii-go/pkg/scanii/models"
)

// Retrieves a previously processed file resource - https://docs.scanii.com/v2.1/resources.html#files
func (c *Client) RetrieveAccountInfo() (*models.AccountInfo, error) {

	req, err := http.NewRequest("GET", endpoints.Resolve(c.Target, "account.json"), nil)
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

	var r models.AccountInfo
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}
