package im_log

import (
	"fmt"
	"path"
	"runtime"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder, // INFO

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey: "caller",
			// EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
				// TODO: consider using a byte-oriented API to save an allocation.
				caller.TrimmedPath()
				enc.AppendString(getCallerInfoForLog())
			},
		},
	}.Build()
	// logger, _ = zap.NewProduction()
}

func getCallerInfoForLog() (callerStr string) {
	pc, file, line, ok := runtime.Caller(7) // 回溯七层，拿到写日志的调用方的函数信息
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名
	callerStr = file + "/" + funcName + ":" + strconv.Itoa(line)
	return
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
