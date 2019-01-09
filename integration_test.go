package scanii

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uvasoftware/scanii-go/endpoints"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

var key = strings.Split(os.Getenv("SCANII_CREDS"), ":")[0]
var secret = strings.Split(os.Getenv("SCANII_CREDS"), ":")[1]

func TestPing(t *testing.T) {
	client := NewClient(&ClientOpts{
		Target: endpoints.LATEST,
		Key:    key,
		Secret: secret,
	})

	response, err := client.Ping()
	require.Nil(t, err)
	require.True(t, response, "ping must pong back")
}

func TestPingWithBadCredentials(t *testing.T) {
	client := NewClient(&ClientOpts{
		Target: endpoints.LATEST,
		Key:    "hello",
		Secret: "world",
	})

	response, err := client.Ping()
	require.NotNil(t, err)
	require.False(t, response)
}

func TestPingWithBadCredentials2(t *testing.T) {
	client := NewClient(&ClientOpts{
		Target: endpoints.LATEST,
		Key:    "",
		Secret: "",
	})

	response, err := client.Ping()
	require.NotNil(t, err)
	require.False(t, response)
}

func TestProcess(t *testing.T) {
	client := NewClient(&ClientOpts{
		Target: endpoints.LATEST,
		Key:    key,
		Secret: secret,
	})

	file, _ := ioutil.TempFile("", "")
	defer func() {
		_ = file.Close()
		_ = os.Remove(file.Name())
	}()

	_, _ = file.Write([]byte("hello world"))

	metadata := map[string]string{
		"m1": "v1",
		"m2": "v2",
	}

	r, err := client.Process(file.Name(), "", metadata)

	require.Nil(t, err)
	if assert.NotNil(t, r) {
		require.False(t, r.ID == "")
		require.Equal(t, "text/plain", r.ContentType)
		require.True(t, r.ContentLength > 0)
		require.NotNil(t, r.CreationDate)
		require.Equal(t, 0, len(r.Findings))

	} else {
		t.FailNow()
	}

	time.Sleep(1 * time.Second)

	// now retrieving

	r2, err := client.Retrieve(r.ID)
	require.Nil(t, err)

	if assert.NotNil(t, r2) {
		require.False(t, r2.ID == "")
		require.Equal(t, "text/plain", r2.ContentType)
		require.True(t, r2.ContentLength > 0)
		require.NotNil(t, r2.CreationDate)
		require.Equal(t, 0, len(r2.Findings))

	} else {
		t.FailNow()
	}

}

func TestProcessAsync(t *testing.T) {
	client := NewClient(&ClientOpts{
		Target: endpoints.LATEST,
		Key:    key,
		Secret: secret,
	})

	file, _ := ioutil.TempFile("", "")
	defer func() {
		_ = file.Close()
		_ = os.Remove(file.Name())
	}()

	_, _ = file.Write([]byte("hello world"))

	r, err := client.ProcessAsync(file.Name(), "https://httpbin.org/post", make(map[string]string, 0))

	require.Nil(t, err)
	if assert.NotNil(t, r) {
		require.False(t, r.ID == "")
	} else {
		t.FailNow()
	}

	time.Sleep(1 * time.Second)

	// now retrieving

	r2, err := client.Retrieve(r.ID)
	require.Nil(t, err)

	if assert.NotNil(t, r2) {
		require.False(t, r2.ID == "")
		require.Equal(t, "text/plain", r2.ContentType)
		require.True(t, r2.ContentLength > 0)
		require.NotNil(t, r2.CreationDate)
		require.Equal(t, 0, len(r2.Findings))

	} else {
		t.FailNow()
	}

}

func TestProcessWithFindings(t *testing.T) {
	client := NewClient(&ClientOpts{
		Target: endpoints.LATEST,
		Key:    key,
		Secret: secret,
	})

	file, _ := ioutil.TempFile("", "")
	defer func() {
		_ = file.Close()
		_ = os.Remove(file.Name())
	}()

	_, _ = file.Write([]byte(`X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*`))

	r, err := client.Process(file.Name(), "", make(map[string]string, 0))

	require.Nil(t, err)
	if assert.NotNil(t, r) {
		require.False(t, r.ID == "")
		require.Equal(t, "text/plain", r.ContentType)
		require.True(t, r.ContentLength > 0)
		require.NotNil(t, r.CreationDate)
		require.Equal(t, 1, len(r.Findings))

	} else {
		t.FailNow()
	}
}

func TestFetchWithFindings(t *testing.T) {
	client := NewClient(&ClientOpts{
		Target: endpoints.V20_EU1,
		Key:    key,
		Secret: secret,
	})

	r, err := client.Fetch("https://scanii.s3.amazonaws.com/eicarcom2.zip", "https://httpbin.org/post", make(map[string]string, 0))

	require.Nil(t, err)
	if assert.NotNil(t, r) {
		require.False(t, r.ID == "")
	} else {
		t.FailNow()
	}

	time.Sleep(1 * time.Second)

	r2, err := client.Retrieve(r.ID)
	require.Nil(t, err)

	if assert.NotNil(t, r2) {
		require.Equal(t, r.ID, r2.ID)
		require.Equal(t, "application/zip", r2.ContentType)
		require.True(t, r2.ContentLength > 0)
		require.NotNil(t, r2.CreationDate)
		require.Equal(t, 1, len(r2.Findings))

	} else {
		t.FailNow()
	}
}

func TestRetrieveAccountInfo(t *testing.T) {
	client := NewClient(&ClientOpts{
		Target: endpoints.V21_AP1,
		Key:    key,
		Secret: secret,
	})

	account, err := client.RetrieveAccountInfo()
	require.Nil(t, err)

	if assert.NotNil(t, account) {
		require.NotEmpty(t, account.Name)
		//require.NotEmpty(t, account.Subscription)
		require.NotEmpty(t, account.BillingEmail)
		require.NotNil(t, account.CreationDate)
		require.NotNil(t, account.ModificationDate)
		require.True(t, account.StartingBalance > 0)
		require.True(t, account.Balance > 0)
		require.True(t, len(account.Users) > 0)
		require.True(t, len(account.Keys) > 0)

	} else {
		t.FailNow()
	}
}
func TestCreateAndUseAuthToken(t *testing.T) {
	client := NewClient(&ClientOpts{
		Target: endpoints.V21_AP1,
		Key:    key,
		Secret: secret,
	})

	token, err := client.CreateAuthToken(30)
	require.Nil(t, err)

	if assert.NotNil(t, token) {
		require.NotEmpty(t, token.ID)
		require.NotNil(t, token.CreationDate)
		require.NotNil(t, token.ExpirationDate)
	} else {
		t.FailNow()
	}

	// now we attempt to retrieve it
	token2, err := client.RetrieveAuthToken(token.ID)
	if assert.NotNil(t, token2) {
		require.NotEmpty(t, token2.ID)
		require.NotNil(t, token2.CreationDate)
		require.NotNil(t, token2.ExpirationDate)
		require.Equal(t, token.ID, token2.ID)
		require.Equal(t, token.ExpirationDate, token2.ExpirationDate)
		require.Equal(t, token.CreationDate, token2.CreationDate)
	} else {
		t.FailNow()
	}

	// now we attempt to use it
	client2 := NewClient(&ClientOpts{
		Target: endpoints.LATEST,
		Key:    token.ID,
	})
	_, err = client2.Ping()
	require.Nil(t, err)

	// now we attempt to delete it
	err = client.DeleteAuthToken(token.ID)
	require.Nil(t, err)

}
