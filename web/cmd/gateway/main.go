package main

import (
	"github.com/zxzixuanwang/go-forum/web/gateway"
	"github.com/zxzixuanwang/go-forum/web/gateway/config"
	"github.com/zxzixuanwang/go-forum/web/pkg/logzap"
)

func main() {
	logzap.LogInit(&logzap.LogConfig{
		Console: logzap.Console{
			Enabled: config.ConfigCollection.LogConfigs.LogConsole.Enabled,
			Level:   config.ConfigCollection.LogConfigs.LogConsole.Level,
			Al:      config.ConfigCollection.LogConfigs.LogConsole.Al,
		},
		File: logzap.File{
			Enabled:    config.ConfigCollection.LogConfigs.LogFile.Enabled,
			Level:      config.ConfigCollection.LogConfigs.LogFile.Level,
			Path:       config.ConfigCollection.LogConfigs.LogFile.Path,
			Name:       config.ConfigCollection.LogConfigs.LogFile.Name,
			MaxHistory: config.ConfigCollection.LogConfigs.LogFile.MaxHistory,
			MaxSizeMb:  config.ConfigCollection.LogConfigs.LogFile.MaxSizeMb,
		},
	})
	gateway.LoadRoute()
}
