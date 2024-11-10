package webapp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	app := NewApp(&Config{Port: 4000, LogLevel: "info"})

	serverErrorResponse := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := fmt.Errorf("Internal Server Error")
		app.ServerErrorResponse(w, r, err)
	})

	badRequestResponse := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := fmt.Errorf("Bad Request")
		app.BadRequestResponse(w, r, err)
	})

	notFoundResponse := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.NotFoundResponse(w, r)
	})

	failedValidationResponse := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.FailedValidationResponse(w, r, map[string]string{"field1": "validation failed"})
	})

	methodNotAllowedResponse := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.MethodNotAllowedResponse(w, r)
	})

	editConflictResponse := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.EditConflictResponse(w, r)
	})

	rateLimitExceededResponse := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.RateLimitExceededResponse(w, r)
	})

	tests := []struct {
		name    string
		errFunc http.HandlerFunc
		status  int
	}{
		{
			name:    "Server Error Response",
			errFunc: serverErrorResponse,
			status:  http.StatusInternalServerError,
		},
		{
			name:    "Bad Request Response",
			errFunc: badRequestResponse,
			status:  http.StatusBadRequest,
		},
		{
			name:    "Not Found Response",
			errFunc: notFoundResponse,
			status:  http.StatusNotFound,
		},
		{
			name:    "Failed Validation Response",
			errFunc: failedValidationResponse,
			status:  http.StatusUnprocessableEntity,
		},
		{
			name:    "Method Not Allowed Response",
			errFunc: methodNotAllowedResponse,
			status:  http.StatusMethodNotAllowed,
		},
		{
			name:    "Edit Conflict Response",
			errFunc: editConflictResponse,
			status:  http.StatusConflict,
		},
		{
			name:    "Rate Limit Exceeded Response",
			errFunc: rateLimitExceededResponse,
			status:  http.StatusTooManyRequests,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			rr.Body.Reset()

			r, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			if err != nil {
				t.Fatal(err)
			}
			tt.errFunc.ServeHTTP(rr, r)
			rs := rr.Result()
			assert.Equal(t, tt.status, rs.StatusCode)
			defer rs.Body.Close()
		})
	}
}
