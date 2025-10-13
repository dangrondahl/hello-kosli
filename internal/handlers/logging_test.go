package handlers

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/dangrondahl/hello-kosli/internal/logging"
)

func TestHandlersEmitLogs(t *testing.T) {
    var buf bytes.Buffer
    SetLogger(logging.NewStdLogger(&buf))

    // Call Hello with query param
    req := httptest.NewRequest(http.MethodPost, "/hello?name=Bob", nil)
    rec := httptest.NewRecorder()
    Hello(rec, req)

    // Call Version
    req2 := httptest.NewRequest(http.MethodGet, "/version", nil)
    rec2 := httptest.NewRecorder()
    Version(rec2, req2)

    out := buf.String()
    if !strings.Contains(out, `"msg":"handled request"`) {
        t.Fatalf("expected log lines to contain handled request messages; got: %q", out)
    }
}
