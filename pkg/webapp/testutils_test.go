package webapp

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestApp(t *testing.T) *App {
	config := &Config{
		Port:     8080,
		LogLevel: "info",
	}

	app := NewApp(config)
	return app
}

type TestServer struct {
	*httptest.Server
}

func NewTestServer(h http.Handler) *TestServer {
	ts := httptest.NewServer(h)

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &TestServer{ts}
}

func (ts *TestServer) Get(t *testing.T, urlPath string, partial bool) (int, http.Header, string) {
	req, err := http.NewRequest(http.MethodGet, ts.URL+urlPath, nil)
	if err != nil {
		t.Fatal(err)
	}
	if partial {
		req.Header.Set("HX-Request", "true")
	}

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = rs.Body.Close()
	}()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}
