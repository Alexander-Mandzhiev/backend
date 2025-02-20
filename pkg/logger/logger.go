package sl

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

const (
	envDevelopment = "development"
	envProduction  = "production"
)

func SetupLogger(env string) {
	var level slog.Level
	switch env {
	case envDevelopment:
		level = slog.LevelDebug
	case envProduction:
		level = slog.LevelInfo
	default:
		level = slog.LevelInfo
	}
	Log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
}
