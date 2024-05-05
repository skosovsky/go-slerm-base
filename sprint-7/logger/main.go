package main

import (
	"context"
	"log/slog"
	"os"
)

type HandlerMiddleware struct {
	next slog.Handler
}

func NewHandlerMiddleware(next slog.Handler) *HandlerMiddleware {
	return &HandlerMiddleware{next: next}
}

func (h *HandlerMiddleware) Enabled(ctx context.Context, rec slog.Level) bool {
	return h.next.Enabled(ctx, rec)
}

func (h *HandlerMiddleware) Handle(ctx context.Context, rec slog.Record) error {
	if logStruct, ok := ctx.Value(keyUserID).(logCtx); ok {
		rec.Add("userID", logStruct.userID)
	}

	return h.next.Handle(ctx, rec) //nolint:wrapcheck // middleware
}

func (h *HandlerMiddleware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &HandlerMiddleware{next: h.next.WithAttrs(attrs)}
}

func (h *HandlerMiddleware) WithGroup(name string) slog.Handler {
	return &HandlerMiddleware{next: h.next.WithGroup(name)}
}

type logCtx struct {
	userID int
}

type keyCtx int

const keyUserID = keyCtx(0)

func WithLoginUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, keyUserID, logCtx{userID: userID})
}

func InitLogger() {
	handler := slog.Handler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}))

	handler = NewHandlerMiddleware(handler)
	slog.SetDefault(slog.New(handler))
}

func TransmitSMS(ctx context.Context, phone, message string) error {
	slog.InfoContext(ctx, "Transmit SMS OK", "phone", phone, "message", message)

	return nil
}

func sendSMS(ctx context.Context, phone string) error {
	slog.InfoContext(ctx, "Sending SMS", "phone", phone)
	message := "Спасибо"
	_ = TransmitSMS(ctx, phone, message)
	slog.InfoContext(ctx, "Sent SMS OK", "phone", phone, "message", message)

	return nil
}

func GetPhoneByID(ctx context.Context, userID int) (string, error) { //nolint:revive // example
	phone := "+78009641020"
	slog.InfoContext(ctx, "Get phone OK", "phone", phone)

	return phone, nil
}

func Handler(ctx context.Context, userID int) {
	ctx = WithLoginUserID(ctx, userID)
	slog.InfoContext(ctx, "Handler started")
	phone, _ := GetPhoneByID(ctx, userID)
	_ = sendSMS(ctx, phone)
	slog.InfoContext(ctx, "Handler finished")
}

func main() {
	InitLogger()

	ctx := context.Background()
	Handler(ctx, 1)
}
