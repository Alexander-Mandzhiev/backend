package sl

import (
	"bytes"
	"io"
	"log/slog"
)

type LoggerWriter struct {
	logger *slog.Logger
}

// Write реализует метод io.Writer
func (w *LoggerWriter) Write(p []byte) (n int, err error) {
	// Преобразуем входные данные в строку и записываем их через slog.Logger
	message := string(bytes.TrimSpace(p))
	w.logger.Info(message)
	return len(p), nil
}

// NewLoggerWriter создает новый экземпляр LoggerWriter
func NewLoggerWriter(logger *slog.Logger) io.Writer {
	return &LoggerWriter{logger: logger}
}
