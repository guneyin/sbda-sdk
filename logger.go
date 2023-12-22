package sdk

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger() *Logger {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

	return &Logger{logger: logger}
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	l.logger.WithContext(ctx)

	return l
}

func (l *Logger) AddStrField(k, v string) *Logger {
	l.logger.Log().Str(k, v)

	return l
}

func (l *Logger) AddIntField(k string, v int) *Logger {
	l.logger.Log().Int(k, v)

	return l
}

func (l *Logger) AddBoolField(k string, v bool) *Logger {
	l.logger.Log().Bool(k, v)

	return l
}

func (l *Logger) Debug(log string) {
	l.logger.Debug().Msg(log)
}

func (l *Logger) Warn(log string) {
	l.logger.Warn().Msg(log)
}

func (l *Logger) Info(log string) {
	l.logger.Info().Msg(log)
}

func (l *Logger) Error(log string) {
	l.logger.Error().Msg(log)
}

func (l *Logger) Fatal(log string) {
	l.logger.Fatal().Msg(log)
}
