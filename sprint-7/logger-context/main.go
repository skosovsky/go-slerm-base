package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"os"
)

type ctxKey string

func main() {
	var requestID ctxKey = "request-id"
	ctx := context.WithValue(context.Background(), requestID, "req-123")

	var userID = 1
	if err := CreateUser(ctx, userID); err != nil {
		slog.Error("failed to create user",
			slog.Any("error", err))
	}

	loggerErr := slog.New(slog.NewTextHandler(os.Stderr, nil))
	loggerErr.Info("error")

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	loggerFile := slog.New(slog.NewJSONHandler(file, nil))
	loggerFile.Info("info")
}

func CreateUser(ctx context.Context, userID int) error {
	requestID, _ := ctx.Value(ctxKey("request-id")).(string)
	slogger := slog.With(
		slog.String("request-id", requestID),
		slog.Int("user-id", userID),
	)

	slogger.Info("user creation started")

	if err := SaveUser(ctx, userID); err != nil {
		slogger.Error("user creation failed",
			slog.Any("error", err))

		return err
	}
	slogger.Info("user creation done")

	return nil
}

func SaveUser(_ context.Context, _ int) error {
	return errors.New("not implemented")
}
