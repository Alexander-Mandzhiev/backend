package sl

import (
	"bytes"
	"io"
	"log/slog"
)

type LoggerWriter struct {
	logger *slog.Logger
}

func (w *LoggerWriter) Write(p []byte) (n int, err error) {
	message := string(bytes.TrimSpace(p))
	w.logger.Info(message)
	return len(p), nil
}

func NewLoggerWriter(logger *slog.Logger) io.Writer {
	return &LoggerWriter{logger: logger}
}
