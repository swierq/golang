package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/swierq/golang/internal/loadek"
	"github.com/swierq/golang/pkg/webapp"
)

func newTestApp() *loadekApp {
	webConfig := webapp.NewConfig(webapp.WithPort(uint16(4040)), webapp.WithLogLevel("info"))
	webapp := webapp.NewApp(webConfig)

	loadekConfig, _ := loadek.NewConfig(loadek.WithCpuLoadMi(20), loadek.WithMemLoadMb(100))
	loadek := loadek.NewApp(loadek.WithConfig(loadekConfig))

	app, _ := newApp(withLoadek(loadek), withWebApp(webapp))
	return app

}

type testServer struct {
	*httptest.Server
}

func newTestServer(h http.Handler) *testServer {
	ts := httptest.NewServer(h)

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string, partial bool) (int, http.Header, string) {
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
