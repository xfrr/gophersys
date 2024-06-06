package logger

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

// Logger is a reference to the zerolog.Logger type
type Logger = zerolog.Logger

// Level is a reference to the zerolog.Level type
type Level = zerolog.Level

// NewLogger creates a new zerolog.Logger instance
func NewLogger(svcName string, level Level) Logger {
	buildInfo, _ := debug.ReadBuildInfo()
	zerolog.SetGlobalLevel(level)
	return zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		},
	).With().
		Timestamp().
		Str("component", svcName).
		Str("go_version", buildInfo.GoVersion).
		Logger()
}
