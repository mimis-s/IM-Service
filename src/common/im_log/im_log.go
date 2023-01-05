package im_log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func NewLogger(logPath string) {

	options := []zap.Option{
		zap.AddCallerSkip(1),             // 上一层日志调用路径
		zap.AddStacktrace(zap.WarnLevel), // 设置warn/error级别的日志会输出堆栈调用
	}

	logger, _ = zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder, // INFO

			TimeKey:    "time",
			EncodeTime: zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),

			CallerKey: "caller",
			EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
				// TODO: consider using a byte-oriented API to save an allocation.
				enc.AppendString(trimmedPath(caller))
				// enc.AppendString(caller.TrimmedPath())
			},
		},
	}.Build(options...)
}

func Info(v ...interface{}) {
	info := fmt.Sprintf(v[0].(string), v[1:]...)
	logger.Info(info)
}

func Warn(v ...interface{}) {
	info := fmt.Sprintf(v[0].(string), v[1:]...)
	logger.Warn(info)
}

func Error(v ...interface{}) {
	info := fmt.Sprintf(v[0].(string), v[1:]...)
	logger.Error(info)
}
