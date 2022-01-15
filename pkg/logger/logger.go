package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
)

var Logger *zap.Logger

func Init() (*zap.Logger, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.TimeKey = "date"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if zapLogger, err := zapConfig.Build(); err != nil {
		return nil, err
	} else {
		zapLogger = zapLogger.With(
			zap.String("os", runtime.GOOS),
			zap.String("version", runtime.Version()),
			zap.String("arch", runtime.GOARCH),
		)
		Logger = zapLogger
		return Logger, nil
	}
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}
