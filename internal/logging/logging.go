package logging

import (
	"encoding/json"
	"io"
	"time"
)

// Logger is a tiny structured logger used by the handlers.
type Logger interface {
    Info(msg string, fields map[string]interface{})
    Error(msg string, fields map[string]interface{})
}

type stdLogger struct{
    out io.Writer
}

// NewStdLogger writes JSON logs to the provided writer.
func NewStdLogger(w io.Writer) Logger {
    return &stdLogger{out: w}
}

func (l *stdLogger) Info(msg string, fields map[string]interface{}) {
    l.log("info", msg, fields)
}

func (l *stdLogger) Error(msg string, fields map[string]interface{}) {
    l.log("error", msg, fields)
}

func (l *stdLogger) log(level, msg string, fields map[string]interface{}) {
    m := map[string]interface{}{
        "level": level,
        "msg":   msg,
        "time":  time.Now().Format(time.RFC3339),
    }
    for k, v := range fields {
        m[k] = v
    }
    if b, err := json.Marshal(m); err == nil {
        _, _ = l.out.Write(append(b, '\n'))
    }
}

type nopLogger struct{}

func (n nopLogger) Info(_ string, _ map[string]interface{})  {}
func (n nopLogger) Error(_ string, _ map[string]interface{}) {}

// Nop returns a logger that does nothing. Useful as the default.
func Nop() Logger { return nopLogger{} }
