package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"strings"
	"time"
)

type customJSONHandler struct {
	w    io.Writer
	opts *slog.HandlerOptions
}

func newCustomJSONHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return &customJSONHandler{
		w:    w,
		opts: opts,
	}
}

func (handler *customJSONHandler) Enabled(_ context.Context, logLevel slog.Level) bool {
	if handler.opts != nil && handler.opts.Level != nil {
		return logLevel >= handler.opts.Level.Level()
	}
	return true
}

func (handler *customJSONHandler) Handle(_ context.Context, record slog.Record) error {
	layout := "2006-01-02T15:04:05.999999Z"
	date, err := time.Parse(layout, record.Time.Format(layout))
	if err != nil {
		return err
	}
	data := map[string]any{
		"time":    date,
		"level":   record.Level.String(),
		"message": record.Message,
	}

	record.Attrs(func(attr slog.Attr) bool {
		mapAttributes(attr, data)
		return true
	})

	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if err := json.Indent(&out, raw, "", "  "); err != nil {
		return err
	}
	out.WriteByte('\n')

	_, err = handler.w.Write(out.Bytes())
	return err
}

func mapAttributes(attr slog.Attr, data map[string]any) {
	if attr.Value.Kind() == slog.KindGroup {
		requestData := attr.Value.String()
		requestData = strings.TrimPrefix(requestData, "[")
		requestData = strings.TrimSuffix(requestData, "]")
		requestDataArr := strings.Split(requestData, " ")
		var m = make(map[string]string)
		for param, _ := range requestDataArr {
			paramArr := strings.Split(requestDataArr[param], "=")
			m[paramArr[0]] = paramArr[1]
		}
		data[attr.Key] = m
	} else {
		data[attr.Key] = attr.Value.String()
	}
}

func (handler *customJSONHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return handler
}

func (handler *customJSONHandler) WithGroup(name string) slog.Handler {
	return handler
}
