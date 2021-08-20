package config

import "github.com/spf13/viper"

var ConfigCollection ConfigMap
var DefaultConfig = ConfigMap{
	Services: Service{Env: "test", Port: ":8000", Name: "forum-gateway"},
	LogConfigs: LogConfig{
		LogConsole: Console{Enabled: true, Level: "debug", Al: "info"},
		LogFile:    File{Enabled: false},
	},
}

type ConfigMap struct {
	LogConfigs LogConfig
	Services   Service
}

type LogConfig struct {
	LogConsole Console `json:"console"`
	LogFile    File    `json:"file"`
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
type Service struct {
	Env  string `json:"env"`
	Port string `json:"port"`
	Name string `json:"name"`
}

func init() {
	viper.AddConfigPath("./web/configs")
	viper.SetConfigName("gateway")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config err:" + err.Error())
	}
	ConfigCollection = DefaultConfig
	err = viper.Unmarshal(&ConfigCollection)
	if err != nil {
		panic("get config error:" + err.Error())
	}
}
