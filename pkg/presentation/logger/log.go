package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}).With().Timestamp().Logger(),
	}
}

func (l Logger) Log(level zerolog.Level, message string, format ...interface{}) error {
	switch level {
	case zerolog.ErrorLevel:
		l.logger.Error().Msgf(message, format...)
	case zerolog.DebugLevel:
		l.logger.Debug().Msgf(message, format...)
	case zerolog.FatalLevel:
		l.logger.Fatal().Msgf(message, format...)

	default:
		l.logger.Info().Msgf(message, format...)
	}
	return nil
}
