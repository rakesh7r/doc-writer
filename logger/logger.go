package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func SetupLogger(level string) {
	currentLevel := slog.LevelInfo

	switch level {
	case "debug":
		currentLevel = slog.LevelDebug
	case "info":
		currentLevel = slog.LevelInfo
	case "warn":
		currentLevel = slog.LevelWarn
	case "error":
		currentLevel = slog.LevelError
	}

	options := &slog.HandlerOptions{
		Level: currentLevel,
	}

	handler := slog.NewTextHandler(os.Stdout, options)
	Log = slog.New(handler)
	slog.SetDefault(Log)
}
