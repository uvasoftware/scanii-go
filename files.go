package scaniigo

import (
	"encoding/json"
	"net/http"
)

// ProcessedFileResponse holds the returned value from a call to
// the previously processed file endpoint
type ProcessFileResponse struct {
	ID            string   `json:"id"`
	Checksum      string   `json:"checksum"`
	ContentLength int      `json:"content_length"`
	Findings      []string `json:"findings"`
	CreationDate  string   `json:"creation_date"`
	ContentType   string   `json:"content_type"`
}

// ProcessFileRequest holds; the options needed for the given API call
type ProcessFileParams struct {
	FileLocation string
	Callback     string
	Metadata     string
}

// RetrieveProcessedFile retrieves a previously processed file resource
func (c *Client) RetrieveProcessedFile(id string) (*ProcessFileResponse, error) {
	req, err := http.NewRequest("GET", c.Endpoint+FilePath+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.APIAuth.Key, c.APIAuth.Secret)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pfr ProcessFileResponse
	if err := json.NewDecoder(res.Body).Decode(&pfr); err != nil {
		return nil, err
	}
	return &pfr, nil
}

// ProcessFileSync submits a file for processing synchronously
func (c *Client) ProcessFileSync(pfp *ProcessFileParams) (*ProcessFileResponse, error) {
	req, err := http.NewRequest("POST", c.Endpoint+FilePath, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.APIAuth.Key, c.APIAuth.Secret)
	req.Header.Set("Content-Type", "multipart/form-data")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pfr ProcessFileResponse
	if err := json.NewDecoder(res.Body).Decode(&pfr); err != nil {
		return nil, err
	}
	return &pfr, nil
}

// ProcessFileAsync submits a file for processing synchronously
func (c *Client) ProcessFileAsync(pfp *ProcessFileParams) (*ProcessFileResponse, error) {
	req, err := http.NewRequest("POST", c.Endpoint+FileAsyncPath, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.APIAuth.Key, c.APIAuth.Secret)
	req.Header.Set("Content-Type", "multipart/form-data")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pfr ProcessFileResponse
	if err := json.NewDecoder(res.Body).Decode(&pfr); err != nil {
		return nil, err
	}
	return &pfr, nil
}
