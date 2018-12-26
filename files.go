package scanii

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uvasoftware/scanii-go/endpoints"
	"github.com/uvasoftware/scanii-go/models"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const formParamName = "file"

// Submits a file for processing synchronously - https://docs.scanii.com/v2.1/resources.html#files
func (c *Client) Process(path, callback string, metadata map[string]string) (*models.ProcessingResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(formParamName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	if len(callback) > 0 {
		_ = writer.WriteField("callback", callback)
	}

	for key, val := range metadata {
		_ = writer.WriteField(fmt.Sprintf("metadata[%s]", key), val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoints.Resolve(c.Target, "files"), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(userAgentHeader, c.UserAgentHeader)
	req.Header.Set(authorizationHeader, c.AuthenticationHeader)
	req.Header.Set(contentTypeHeader, writer.FormDataContentType())

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(string(content))
	}

	var r models.ProcessingResult
	if err := json.Unmarshal(content, &r); err != nil {
		return nil, err
	}
	return &r, nil

}

// Submits a file for processing asynchronously - https://docs.scanii.com/v2.1/resources.html#files
func (c *Client) ProcessAsync(path, callback string, metadata map[string]string) (*models.PendingResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(formParamName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	if len(callback) > 0 {
		_ = writer.WriteField("callback", callback)
	}

	for key, val := range metadata {
		_ = writer.WriteField(fmt.Sprintf("metadata[%s]", key), val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoints.Resolve(c.Target, "files/async"), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(userAgentHeader, c.UserAgentHeader)
	req.Header.Set(authorizationHeader, c.AuthenticationHeader)
	req.Header.Set(contentTypeHeader, writer.FormDataContentType())

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusAccepted {
		return nil, errors.New(string(content))
	}

	var r models.PendingResult
	if err := json.Unmarshal(content, &r); err != nil {
		return nil, err
	}
	return &r, nil

}

// Retrieves a previously processed file resource - https://docs.scanii.com/v2.1/resources.html#files
func (c *Client) Retrieve(id string) (*models.ProcessingResult, error) {

	req, err := http.NewRequest("GET", endpoints.Resolve(c.Target, "files/")+id, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(userAgentHeader, c.UserAgentHeader)
	req.Header.Set(authorizationHeader, c.AuthenticationHeader)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(content))
	}

	var r models.ProcessingResult
	if err := json.Unmarshal(content, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// Submits a remote file to be processed asynchronously https://docs.scanii.com/v2.1/resources.html#files
func (c *Client) Fetch(location, callback string, metadata map[string]string) (*models.PendingResult, error) {

	body := url.Values{}
	body.Set("location", location)

	if len(callback) > 0 {
		body.Set("callback", callback)
	}

	for key, val := range metadata {
		body.Set(fmt.Sprintf("metadata[%s]", key), val)
	}

	req, err := http.NewRequest("POST", endpoints.Resolve(c.Target, "files/fetch"), strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set(userAgentHeader, c.UserAgentHeader)
	req.Header.Set(authorizationHeader, c.AuthenticationHeader)
	req.Header.Set(contentTypeHeader, "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusAccepted {
		return nil, errors.New(string(content))
	}

	var r models.PendingResult
	if err := json.Unmarshal(content, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
