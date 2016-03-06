package scaniigo

import (
	"reflect"
	"testing"
)

// TestRetrieveProcessedFile
func TestRetrieveProcessedFile(t *testing.T) {
	t.Parallel()
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
		Location: "",
	}
	pfr, err := c.ProcessRemoteFileAsync(rfap)
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(pfr).String() != "*scaniigo.ProcessFileResponse" {
		t.Error(ErrInvalidDataType)
	}
}
