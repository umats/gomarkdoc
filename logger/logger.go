// Package logger provides a simple console logger for reporting information
// about execution to stderr.
//
//nolint:govet // This legacy implementation intentionally relies on these patterns.
package logger

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type (
	// Logger provides basic logging capabilities at different logging levels.
	Logger interface {
		Debug(a ...any)
		Debugf(format string, a ...any)
		Info(a ...any)
		Infof(format string, a ...any)
		Warn(a ...any)
		Warnf(format string, a ...any)
		Error(a ...any)
		Errorf(format string, a ...any)
	}

	// Level defines valid logging levels for a Logger.
	Level int

	// Option defines an option for configuring the logger.
	Option func(opts *options)

	// options defines options for configuring the logger.
	options struct {
		fields map[string]any
	}
)

// Valid logging levels.
const (
	DebugLevel Level = iota + 1
	InfoLevel
	WarnLevel
	ErrorLevel
)

// New initializes a new Logger.
func New(level Level, opts ...Option) Logger {
	var options options
	for _, opt := range opts {
		opt(&options)
	}

	log := logrus.New()

	formatter := &prefixed.TextFormatter{
		DisableTimestamp: true,
	}
	formatter.SetColorScheme(&prefixed.ColorScheme{
		DebugLevelStyle: "cyan",
		PrefixStyle:     "black+h",
	})

	log.Formatter = formatter

	switch level {
	case DebugLevel:
		log.SetLevel(logrus.DebugLevel)
	case InfoLevel:
		log.SetLevel(logrus.InfoLevel)
	case WarnLevel:
		log.SetLevel(logrus.WarnLevel)
	case ErrorLevel:
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.ErrorLevel)
	}

	if options.fields != nil {
		return log.WithFields(options.fields)
	}

	return log
}

// WithField sets the provided key/value pair for use on all logs.
func WithField(key string, value any) Option {
	return func(opts *options) {
		if opts.fields == nil {
			opts.fields = make(map[string]any)
		}

		opts.fields[key] = value
	}
}
