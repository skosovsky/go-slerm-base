package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func init() { //nolint:gochecknoinits // it's learning code
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{}) //nolint:exhaustruct // it's good

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see bellow for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.InfoLevel)
}

func loggingExamples() {
	log.Println("test") // 2024/04/26 06:40:32 test

	slog.Info("test", "param", "value") // 2024/04/26 06:41:34 INFO test param=value

	logger := slog.Default()
	logger.Info("test with default logger", "param", "value") // 2024/04/26 06:42:26 INFO test with default logger param=value

	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("test with default logger", "param", "value") // time=2024-04-26T06:43:40.871+03:00 level=INFO msg="test with default logger" param=value

	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("test with default logger", "param", "value") // {"time":"2024-04-26T06:44:22.681921+03:00","level":"INFO","msg":"test with default logger","param":"value"}

	logrus.WithFields(logrus.Fields{
		"param": "value",
	}).Info("Log string from logrus") // {"level":"info","msg":"Log string from logrus","param":"value","time":"2024-04-26T06:46:44+03:00"}

	// A common pattern is to re-use fields between logging statements by re-using the logrus.
	// Entry returned from WithFields() creates copy of struct internally
	contextLogger := logrus.WithFields(logrus.Fields{
		"param":   "value",
		"userID":  12345, //nolint:mnd // it's example
		"traceID": uuid.New().String(),
	})

	contextLogger.Info("Log with fields")
	// {"level":"info","msg":"Log with fields","param":"value","time":"2024-04-26T06:50:09+03:00","traceID":"05d13ebd-b6f2-4bd7-a0f6-1b4402c3c978","userID":12345}

	contextLogger.WithField("hello", "world").Info("Log with fields")
	// {"hello":"world","level":"info","msg":"Log with fields","param":"value","time":"2024-04-26T06:51:59+03:00","traceID":"12350601-e574-453b-afb1-d3177904db09","userID":12345}
}

func main() {
	loggingExamples()
}
