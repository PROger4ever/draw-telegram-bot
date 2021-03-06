package config

import (
	"strings"

	"github.com/PROger4ever-Golang/draw-telegram-bot/error"
	"github.com/jinzhu/configor"
)

const loadingFailed = "Loading config failed"

type BotApiConfig struct {
	ID                  int    `required:"true"`
	Key                 string `required:"true"`
	Debug               bool
	DisableNotification bool   `env:"BOTAPI_DISABLE_NOTIFICATION"`
	ProxyUrl            string `env:"BOTAPI_PROXY_URL"`
	ProxyTimeout        int64  `required:"true" env:"BOTAPI_PROXY_TIMEOUT"`
}

type Config struct {
	BotApi  BotApiConfig
	UserApi struct {
		Host      string `required:"true"`
		Port      int    `required:"true"`
		PublicKey string `required:"true"`
		ApiId     int    `required:"true"`
		ApiHash   string `required:"true"`
		Debug     int
	}
	Management struct {
		OwnerUsername string `required:"true" env:"MANAGEMENT_OWNER_USERNAME"`
		ChatUsername  string `required:"true" env:"MANAGEMENT_CHAT_USERNAME"`
	}
	Mongo struct {
		Host string `required:"true"`
		Port int    `required:"true"`
	}
}

func LoadConfig(file string) (*Config, *eepkg.ExtendedError) {
	var config Config
	errStd := configor.New(&configor.Config{ENVPrefix: "-"}).Load(&config, file)
	if errStd != nil {
		return &config, eepkg.Wrap(errStd, false, true, loadingFailed)
	}
	config.UserApi.PublicKey = strings.Replace(config.UserApi.PublicKey, "\\n", "\n", -1)
	return &config, nil
}
