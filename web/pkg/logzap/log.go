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
	console console
	file    file
}

type console struct {
	Enabled bool   `json:"enabled"`
	Level   string `json:"level"`
	Al      string `json:"al"`
}

type file struct {
	Enabled    bool   `json:"enabled"`
	Level      string `json:"level"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	MaxHistory int    `json:"maxHistory"`
	MaxSizeMb  int    `json:"maxSizeMb"`
}

func LogInit(lc *LogConfig) {

	var zapcores []zapcore.Core
	if lc.console.Enabled {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleLevel := new(zapcore.Level)
		if err := consoleLevel.UnmarshalText([]byte(lc.console.Level)); err != nil {
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
	if lc.file.Enabled {
		fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		fileLevel := new(zapcore.Level)
		if err := fileLevel.UnmarshalText([]byte(lc.file.Level)); err != nil {
			panic(err)
		}
		w := &lumberjack.Logger{
			Filename:   lc.file.Name,
			MaxSize:    lc.file.MaxSizeMb,
			MaxBackups: lc.file.MaxHistory,
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
