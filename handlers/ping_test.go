package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedCode   int
		expectedBody   string
	}{
		{
			name:           "GET request returns Pong",
			method:         http.MethodGet,
			expectedCode:   http.StatusOK,
			expectedBody:   "Pong!",
		},
		{
			name:           "POST request returns method not allowed",
			method:         http.MethodPost,
			expectedCode:   http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/v1/ping", nil)
			w := httptest.NewRecorder()

			PingHandler(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			if w.Body.String() != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, w.Body.String())
			}
		})
	}
}