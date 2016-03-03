package scaniigo

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

// ProcessFileResponse holds the returned value from a call to
// the previously processed file endpoint
type ProcessFileResponse struct {
	ID            string   `json:"id"`
	Checksum      string   `json:"checksum"`
	ContentLength int      `json:"content_length"`
	Findings      []string `json:"findings"`
	CreationDate  string   `json:"creation_date"`
	ContentType   string   `json:"content_type"`
}

// AsyncFileProcessResponse
type AsyncFileProcessResponse struct {
	ID string `json:"id"`
}

// ProcessFileParams holds the options needed for processing calls
type ProcessFileParams struct {
	// File has the contents of the file to be processed
	File string

	// Callback is an optional callback URL to be notified once processing
	// is completed
	Callback string

	// Metadata is an optional metadata argument to be stored with the resource
	Metadata string
}

// Validate makes sure that the required parameters are present when this time is
// to be used
func (p *ProcessFileParams) Validate() error {
	if p.File == "" {
		return ErrFileFieldEmpty
	}
	return nil
}

// RemoteFileAsyncParams holds the options needed for remote async process calls
type RemoteFileAsyncParams struct {
	// Location contains the URL of the file to be fetched and processed.  This
	// should be escaped prior to processing
	Location string

	// Callback is an optional callback URL to be notified once processing
	// is completed
	Callback string

	// Metadata is an optional metadata argument to be stored with the resource
	Metadata string
}

// Validate makes sure that the required parameters are present when this time is
// to be used
func (r *RemoteFileAsyncParams) Validate() error {
	if r.Location == "" {
		return ErrLocationFieldEmpty
	}
	return nil
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

// builds a bytes buffer with the given file
func (pfp *ProcessFileParams) Generate(c *Client, execType string) (*http.Request, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fh, err := os.Open(pfp.File)
	if err != nil {
		return nil, err
	}
	fw, err := w.CreateFormFile("file", pfp.File)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, fh); err != nil {
		return nil, err
	}

	if fw, err = w.CreateFormField("callback"); err != nil {
		return nil, err
	}
	if _, err = fw.Write([]byte(pfp.Callback)); err != nil {
		return nil, err
	}
	if fw, err = w.CreateFormField("metadata"); err != nil {
		return nil, err
	}
	if _, err = fw.Write([]byte(pfp.Metadata)); err != nil {
		return nil, err
	}
	w.Close()

	var req *http.Request
	switch execType {
	case "async":
		req, err = http.NewRequest("POST", c.Endpoint+FilePath, &b)
		if err != nil {
			return nil, err
		}
	case "sync":
		req, err = http.NewRequest("POST", c.Endpoint+FilePath, &b)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrUnrecognizedExecType
	}

	req.SetBasicAuth(c.APIAuth.Key, c.APIAuth.Secret)
	req.Header.Set("Content-Type", w.FormDataContentType())
	return req, nil
}

// ProcessFileSync submits a file for processing synchronously
func (c *Client) ProcessFileSync(pfp *ProcessFileParams) (*ProcessFileResponse, error) {
	if err := Validate(pfp); err != nil {
		return nil, err
	}

	req, err := GenerateAPIRequest(c, pfp, "sync")
	if err != nil {
		log.Fatalln(err)
	}

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
func (c *Client) ProcessFileAsync(pfp *ProcessFileParams) (*AsyncFileProcessResponse, error) {
	if err := Validate(pfp); err != nil {
		return nil, err
	}

	req, err := GenerateAPIRequest(c, pfp, "async")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var afpr AsyncFileProcessResponse
	if err := json.NewDecoder(res.Body).Decode(&afpr); err != nil {
		return nil, err
	}
	return &afpr, nil
}

// ProcessRemoteFileAsync submits a file for processing synchronously
func (c *Client) ProcessRemoteFileAsync(rfap *RemoteFileAsyncParams) (*AsyncFileProcessResponse, error) {
	if err := Validate(rfap); err != nil {
		return nil, err
	}

	var data url.Values
	data.Set("location", rfap.Location)
	data.Add("callback", rfap.Callback)
	data.Add("metadata", rfap.Metadata)

	req, err := http.NewRequest("POST", c.Endpoint+FileAsyncPath, bytes.NewBufferString(data.Encode()))
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

	var afpr AsyncFileProcessResponse
	if err := json.NewDecoder(res.Body).Decode(&afpr); err != nil {
		return nil, err
	}
	return &afpr, nil
}
