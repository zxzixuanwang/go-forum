package logzap

import (
	"os"
	"sync/atomic"
	"unsafe"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var rootLogger = unsafe.Pointer(zap.NewNop().Sugar())

func GetLogger() *zap.SugaredLogger {
	return (*zap.SugaredLogger)(atomic.LoadPointer(&rootLogger))
}

type LogConfig struct {
	Console Console
	File    File
}

type Console struct {
	Enabled bool   `json:"enabled"`
	Level   string `json:"level"`
	Al      string `json:"al"`
}

type File struct {
	Enabled    bool   `json:"enabled"`
	Level      string `json:"level"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	MaxHistory int    `json:"maxHistory"`
	MaxSizeMb  int    `json:"maxSizeMb"`
}

func LogInit(lc *LogConfig) {

	var zapcores []zapcore.Core
	if lc.Console.Enabled {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleLevel := new(zapcore.Level)
		if err := consoleLevel.UnmarshalText([]byte(lc.Console.Level)); err != nil {
			panic(err)
		}
		zapcores = append(
			zapcores,
			zapcore.NewCore(
				consoleEncoder,
				zapcore.Lock(os.Stdout),
				zap.NewAtomicLevelAt(*consoleLevel),
			),
		)
	}
	if lc.File.Enabled {
		fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		fileLevel := new(zapcore.Level)
		if err := fileLevel.UnmarshalText([]byte(lc.File.Level)); err != nil {
			panic(err)
		}
		w := &lumberjack.Logger{
			Filename:   lc.File.Name,
			MaxSize:    lc.File.MaxSizeMb,
			MaxBackups: lc.File.MaxHistory,
		}
		zapcores = append(
			zapcores,
			zapcore.NewCore(
				fileEncoder,
				zapcore.AddSync(w),
				zap.NewAtomicLevelAt(*fileLevel),
			),
		)
	}
	atomic.StorePointer(
		&rootLogger,
		unsafe.Pointer(zap.New(
			zapcore.NewTee(zapcores...),
			zap.AddStacktrace(zapcore.ErrorLevel),
		).Sugar()),
	)
}
