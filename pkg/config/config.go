package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewConfig(path string) (*viper.Viper, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return viper.GetViper(), nil
}

var ProviderSet = wire.NewSet(NewConfig)