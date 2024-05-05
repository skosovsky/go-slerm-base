package main

import (
	"log/slog"
)

func main() {
	slog.Info("OK (function)", "userID", 111) //nolint:mnd // example

	logger := slog.Default()

	logger = logger.With("data_center", "SomeDC")
	logger.Info("OK (method)", "userID", 222) //nolint:mnd // example
}
