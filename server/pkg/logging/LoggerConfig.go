package logging

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func LoggerSetup() {
	logLevelStr := os.Getenv("LOG_LEVEL")
	logLevel, _ := getLogLevel(logLevelStr)

	environment := os.Getenv("GO_ENV")
	var logger *slog.Logger
	if environment == "development" {
		logger = slog.New(newCustomJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     logLevel,
		}))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			AddSource: true,
			Level:     logLevel,
		}))
	}

	slog.SetDefault(logger)
}

func getLogLevel(logLevelStr string) (logLevel slog.Level, err error) {
	switch strings.ToUpper(logLevelStr) {
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN", "WARNING":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("unknown log level: %s", logLevelStr)
	}
}
