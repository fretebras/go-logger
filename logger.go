package logger

import (
	"github.com/fretebras/go-logger/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const callerFieldName = "trace"

type Logger struct {
	env     string
	app     string
	version string
	logger  zerolog.Logger
	service string
}

var c *config.Config

func New() *Logger {
	c = config.NewConfig()

	l := &Logger{
		env:     c.Environment,
		app:     c.App,
		version: c.Version,
	}

	zerolog.CallerFieldName = callerFieldName
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro

	l.logger = log.With().
		Str("app", l.app).
		Str("version", l.version).
		Str("environment", l.env).
		Logger()
	return l
}

func (l *Logger) Debug(Message string, Context ...string) {
	l.logger.Debug().Str("service", l.service).Strs("context", Context).Msg(Message)
}

func (l *Logger) Info(Message string, Context ...string) {
	l.logger.Info().
		Str("service", l.service).Strs("context", Context).Msg(Message)
}

func (l *Logger) Warning(Message string, Context ...string) {
	l.logger.Warn().
		Str("service", l.service).Strs("context", Context).Msg(Message)
}

func (l *Logger) Error(Error error, Context ...string) {
	l.logger.Error().Str("service", l.service).Strs(
		"context", Context).
		Caller(1).Msg(Error.Error())
}

func (l *Logger) Critical(Error error, Context ...string) {
	l.logger.Fatal().Str("service", l.service).Strs(
		"context", Context).
		Caller(1).Msg(Error.Error())
}

func (l *Logger) SetLogService(service string) {
	l.service = service
}
