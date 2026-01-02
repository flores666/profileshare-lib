package logger

import "log/slog"

func Error(err error) slog.Attr {
	if err == nil {
		return slog.Attr{
			Key:   "error",
			Value: slog.StringValue("unknown error"),
		}
	}

	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
