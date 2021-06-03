package config

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func NewConfig(path string) (*viper.Viper, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "初始化配置文件失败")
	}
	return viper.GetViper(), nil
}

var ConfigSet = wire.NewSet(NewConfig)