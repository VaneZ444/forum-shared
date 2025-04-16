package logger

import (
	"log/slog"
	"os"
)

// New создает экземпляр логгера с настройками уровня.
func New(level slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	return slog.New(handler)
}

// Err добавляет поле с ошибкой в лог.
func Err(err error) slog.Attr {
	return slog.String("error", err.Error())
}
