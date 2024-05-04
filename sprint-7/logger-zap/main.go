package main

import "go.uber.org/zap"

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync() //nolint:errcheck // it's zap

	logger.Info("Hello Zap!")

	logger.Info("Hello Zap again!",
		zap.String("foo", "bar"),
		zap.Int("baz", 1),
		zap.Bool("bool", false),
	)

	sugar := logger.Sugar()
	sugar.Info("Hello Sugar!")
	sugar.Info()

	loggerDev := zap.Must(zap.NewDevelopment())
	defer loggerDev.Sync() //nolint:errcheck // it's zap

	loggerDev.Info("Hello Developer!")
}
