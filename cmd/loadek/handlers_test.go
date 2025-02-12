package main

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swierq/golang/internal/loadek"
)

var (
	tapp *loadekApp
	ts   *testServer
)

func setup() {
	tapp = newTestApp()
	ts = newTestServer(tapp.webapp.Router)
}

func shutdown() {
	ts.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestConfigHandler(t *testing.T) {
	code, _, body := ts.get(t, "/api/config", false)
	assert.Equal(t, http.StatusOK, code)

	var decoded map[string]*loadek.Config
	err := json.Unmarshal([]byte(body), &decoded)
	assert.Nil(t, err)
	assert.Equal(t, tapp.loadek.Config, decoded["config"])

	code, _, body = ts.get(t, "/api/config", true)
	assert.Equal(t, http.StatusOK, code)
	assert.Contains(t, body, "<div")
	err = json.Unmarshal([]byte(body), &decoded)
	assert.NotNil(t, err)
}

func TestUiHandler(t *testing.T) {
	code, _, body := ts.get(t, "/", false)
	assert.Equal(t, http.StatusOK, code)
	assert.Contains(t, body, "</body>")
}
