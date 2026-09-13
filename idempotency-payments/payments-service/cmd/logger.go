package main

import (
	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.uber.org/zap"
)

func InitZapLogger() *zap.Logger {
	core := otelzap.NewCore(
		"github.com/prasannakumar414/system-implementations",
	)

	logger := zap.New(core)
	return logger
}
