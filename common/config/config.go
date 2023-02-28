package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Settings Settings `yaml:"settings"`
}
type Settings struct {
	Database    Database    `yaml:"database"`
	Application Application `yaml:"application"`
	Logger      Logger      `yaml:"logger"`
}
type Database struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}
type Logger struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}
type Application struct {
	Env  string `yaml:"env"`
	Host string `yaml:"host"`
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type ConfigViper struct {
	ConfigViper *viper.Viper
	Config      Config
}

var App = new(ConfigViper)

// 配置文件初始化
func InitConfig() *viper.Viper {
	configFile := "config/config.yml"

	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed:%s", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&App.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&App.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
