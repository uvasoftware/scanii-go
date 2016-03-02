package scaniigo

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
	return nil, nil
}

// ProcessFileSync submits a file for processing synchronously
func (c *Client) ProcessFileSync(pfr *ProcessFileParams) (*ProcessFileResponse, error) {
	return nil, nil
}

// ProcessFileAsync submits a file for processing synchronously
func (c *Client) ProcessFileAsync(pfr *ProcessFileParams) (*ProcessFileResponse, error) {
	return nil, nil
}
