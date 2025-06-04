package logger

import (
	"log/slog"
	"os"
)

type loggerImpl struct {
	l *slog.Logger
}

func New(level slog.Level) Logger {
	opts := &slog.HandlerOptions{
		Level: level,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	return &loggerImpl{
		l: slog.New(handler),
	}
}

func (l *loggerImpl) Debug(msg string, args ...slog.Attr) {
	l.l.Debug(msg, attrsToAny(args)...)
}

func (l *loggerImpl) Info(msg string, args ...slog.Attr) {
	l.l.Info(msg, attrsToAny(args)...)
}

func (l *loggerImpl) Error(msg string, args ...slog.Attr) {
	l.l.Error(msg, attrsToAny(args)...)
}

func attrsToAny(attrs []slog.Attr) []any {
	res := make([]any, len(attrs))
	for i, a := range attrs {
		res[i] = a
	}
	return res
}

func Err(err error) slog.Attr {
	return slog.String("error", err.Error())
}
