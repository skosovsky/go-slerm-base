package main

import (
	"log"
	"log/slog"
	"os"
)

const (
	LevelTrace = slog.Level(8)
	LevelFatal = slog.Level(12)
)

func main() {
	slog.Info("test info")   // INFO test info
	slog.Error("test err")   // ERROR test err
	slog.Warn("test warn")   // WARN test warn
	slog.Debug("test debug") //

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("test info") // {"time":"2024-05-02T15:50:52.426306+03:00","level":"INFO","msg":"test info"}

	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("test info") // time=2024-05-02T15:51:19.977+03:00 level=INFO msg="test info"

	slog.SetDefault(logger)

	slog.Info("test info")   // time=2024-05-02T15:52:21.766+03:00 level=INFO msg="test info"
	log.Println("test info") // time=2024-05-02T15:52:21.766+03:00 level=INFO msg="test info"

	logger.Info("test info", //nolint:govet // time=2024-05-02T15:54:01.846+03:00 level=INFO msg="test info" test=info
		"test", "info",
		"bad") // !BADKEY=bad // linter work

	logger.Info("test info",
		slog.String("test", "string"),
		slog.Int("test", 1),
	)
}
