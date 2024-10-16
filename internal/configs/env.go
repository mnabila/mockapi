package configs

import (
	"fmt"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type EnvConfig struct {
	ServerAddr string `koanf:"SERVER_ADDR"`
	ServerPort int    `koanf:"SERVER_PORT"`

	FinpayNotifyUrl string `koanf:"FINPAY_NOTIFY_URL"`
	OYINotifyUrl    string `koanf:"OYI_NOTIFY_URL"`
	IRSNotifyUrl    string `koanf:"IRS_NOTIFY_URL"`
}

func (c EnvConfig) GetListenAddress() string {
	return fmt.Sprintf("%s:%d", c.ServerAddr, c.ServerPort)
}

func NewEnvConfig() (*EnvConfig, error) {
	conf := koanf.New(".")
	if err := conf.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := new(EnvConfig)

	if err := conf.UnmarshalWithConf("", &config, koanf.UnmarshalConf{FlatPaths: false}); err != nil {
		return nil, fmt.Errorf("failed to read .env file: %v", err)
	}

	return config, nil
}
