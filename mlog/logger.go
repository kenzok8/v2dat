package mlog

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = mustInitLogger()

func mustInitLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = customTimeEncoder
	l, err := config.Build(zap.WithCaller(false))
	if err != nil {
		panic("failed to init mlog:" + err.Error())
	}
	return l
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func L() *zap.Logger {
	return logger
}
