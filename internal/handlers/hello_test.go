package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		query          string
		wantStatus     int
		wantBodyPrefix string
	}{
		{
			name:           "json body",
			method:         http.MethodPost,
			body:           `{"name":"Dan"}`,
			wantStatus:     http.StatusOK,
			wantBodyPrefix: "Hello, Dan",
		},
		{
			name:           "query param",
			method:         http.MethodPost,
			query:          "?name=Alice",
			wantStatus:     http.StatusOK,
			wantBodyPrefix: "Hello, Alice",
		},
		{
			name:           "missing name -> 400",
			method:         http.MethodPost,
			body:           `{}`,
			wantStatus:     http.StatusBadRequest,
			wantBodyPrefix: "missing name",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/hello"+tc.query, bytes.NewBufferString(tc.body))
			rec := httptest.NewRecorder()

			Hello(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.wantStatus {
				t.Fatalf("status: got %d, want %d", res.StatusCode, tc.wantStatus)
			}
			got := rec.Body.String()
			if tc.wantStatus == http.StatusOK && got != tc.wantBodyPrefix {
				t.Fatalf("body: got %q, want %q", got, tc.wantBodyPrefix)
			}
			if tc.wantStatus != http.StatusOK && !bytes.Contains(rec.Body.Bytes(), []byte(tc.wantBodyPrefix)) {
				t.Fatalf("error body should contain %q, got %q", tc.wantBodyPrefix, got)
			}
		})
	}
}
