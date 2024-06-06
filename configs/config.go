package configs

import (
	"github.com/spf13/viper"
)

var v *viper.Viper
var conf *Config

func init() {
	v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs/")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
}

type Config struct {
	AppID     uint64
	AppToken  string
	AppSecret string
}

func NewConfig() *Config {
	conf = &Config{
		AppID:     v.GetUint64("APP_ID"),
		AppToken:  v.GetString("APP_TOKEN"),
		AppSecret: v.GetString("APP_SECRET"),
	}
	return conf
}

func GetConfig() *Config {
	if conf == nil {
		conf = NewConfig()
	}
	return conf
}
