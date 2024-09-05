package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadEnvironment(configPath string) {
	viper.AddConfigPath(configPath)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}

	viper.AutomaticEnv()
}
