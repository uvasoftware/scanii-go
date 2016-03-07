package scaniigo

import (
	"reflect"
	"testing"
)

var localFile = "scan_files/test_binary"
var remoteFile = "https://github.com/uvasoftware/scanii-go/scan_files/test_binary"

// TestRetrieveProcessedFile
func TestRetrieveProcessedFile(t *testing.T) {
	t.Parallel()
}

// TestValidate
func TestValidate(t *testing.T) {
	t.Parallel()
	pfp := &ProcessFileParams{
		File: localFile,
	}
	rfap := &RemoteFileAsyncParams{
		Location: remoteFile,
	}
	var testParams []Validator
	testParams = append(testParams, pfp, rfap)
	for _, i := range testParams {
		if err := Validate(i); err != nil {
			t.Error(err)
		}
	}
}

// TestProcessFileSync
func TestProcessFileSync(t *testing.T) {
	t.Parallel()
	clientOpts := &ClientOpts{
		Version:  "2.1",
		Validate: false,
	}
	c, err := NewClient(clientOpts)
	if err != nil {
		t.Error(err)
	}
	pfp := &ProcessFileParams{
		File: "scan_files/test_binary",
	}
	pfr, err := c.ProcessFileSync(pfp)
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(pfr).String() != "*scaniigo.ProcessFileResponse" {
		t.Error(ErrInvalidDataType)
	}
}

// TestProcessFileAsync
func TestProcessFileAsync(t *testing.T) {
	t.Parallel()
	clientOpts := &ClientOpts{
		Version:  "2.1",
		Validate: false,
	}
	c, err := NewClient(clientOpts)
	if err != nil {
		t.Error(err)
	}
	pfp := &ProcessFileParams{
		File: localFile,
	}
	pfr, err := c.ProcessFileSync(pfp)
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(pfr).String() != "*scaniigo.ProcessFileResponse" {
		t.Error(ErrInvalidDataType)
	}
}

// TestProcessRemoteFileAsync
func TestProcessRemoteFileAsync(t *testing.T) {
	t.Parallel()
	clientOpts := &ClientOpts{
		Version:  "2.1",
		Validate: false,
	}
	c, err := NewClient(clientOpts)
	if err != nil {
		t.Error(err)
	}
	rfap := &RemoteFileAsyncParams{
		Location: remoteFile,
	}
	pfr, err := c.ProcessRemoteFileAsync(rfap)
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(pfr).String() != "*scaniigo.AsyncFileProcessResponse" {
		t.Error(ErrInvalidDataType)
	}
}
