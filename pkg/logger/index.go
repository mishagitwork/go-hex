package logger

import (
	"fmt"
	"os"
	"reflect"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func New() *Logger {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano
	_ = os.Mkdir("./logs", 0750)

	logger := zerolog.
		New(createLoggerWriter()).
		With().Timestamp().
		CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + 1).
		Logger()

	l := &Logger{logger: &logger}

	return l
}

func (l *Logger) Debug(message any) {
	l.logger.Debug().Msg(formatLoggerMessage(message))
}

func (l *Logger) Info(message any) {
	l.logger.Info().Msg(formatLoggerMessage(message))
}

func (l *Logger) Warn(message any) {
	l.logger.Warn().Msg(formatLoggerMessage(message))
}

func (l *Logger) Error(message any) {
	l.logger.Error().Msg(formatLoggerMessage(message))
}

func (l *Logger) Fatal(message any) {
	l.logger.Fatal().Msg(formatLoggerMessage(message))
}

func formatLoggerMessage(message any) string {
	switch parsed := message.(type) {
	case error:
		return parsed.Error()
	case string:
		return parsed
	default:
		return fmt.Sprintf("Message of type %v:\n%+v", reflect.TypeOf(parsed), message)
	}
}
