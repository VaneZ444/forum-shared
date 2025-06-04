package logger

import "log/slog"

type Logger interface {
	Debug(msg string, args ...slog.Attr)
	Info(msg string, args ...slog.Attr)
	Error(msg string, args ...slog.Attr)
}
