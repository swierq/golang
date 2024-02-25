package premierleague

import (
	"bytes"
	"fmt"
	// "github.com/jarcoal/httpmock"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestDummy(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		fmt.Println(req.URL.Path)
		var data []byte
		if req.URL.Path == "/api/fixtures/" {
			data, _ = os.ReadFile("testdata/fixtures.json")
		} else if req.URL.Path == "/api/bootstrap-static/" {
			data, _ = os.ReadFile("testdata/bootstrap-static.json")
		} else {
			data = make([]byte, 0)
		}
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(data)),
			Header:     make(http.Header),
		}
	})
	err := AllFixtures("", client)
	require.NoError(t, err, "There should be no error.")
	result := 1
	require.Equal(t, 1, result, "Result should be 1.")
}
