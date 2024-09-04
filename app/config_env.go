package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadEnvironment(mode string) {
	if mode == "test" {
		viper.AddConfigPath("..")
	} else {
		viper.AddConfigPath(".")
	}
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}

	viper.AutomaticEnv()
}
