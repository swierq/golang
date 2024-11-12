package webapp

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {

	app := newTestApp(t)

	ts := newTestServer(t, app.GetRouter())
	defer ts.Close()

	code, _, body := ts.get(t, "/health", false)

	assert.Equal(t, http.StatusOK, code)
	assert.Contains(t, body, "available")
}
