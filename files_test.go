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
	clientOpts := &ClientOpts{
		Version:  "2.1",
		Validate: false,
	}
	c, err := NewClient(clientOpts)
	if err != nil {
		t.Error(err)
	}
	values := []string{"asdfasdfasdfasdf", ""}
	for _, v := range values {
		pfr, err := c.RetrieveProcessedFile(v)
		if err != nil {
			t.Error(err)
		}
		if reflect.TypeOf(pfr).String() != "*scaniigo.ProcessFileResponse" {
			t.Error(ErrInvalidDataType)
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
	pfr, err := c.ProcessFileAsync(pfp)
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(pfr).String() != "*scaniigo.AsyncFileProcessResponse" {
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
