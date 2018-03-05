package main

import (
	"os"

	"github.com/spf13/viper"
)

const defaultConfPath = "conf/prod.yaml"

func initConfig(filepath string) (err error) {
	if filepath == "" {
		filepath = defaultConfPath
	}
	if _, err = os.Stat(filepath); err != nil {
		return
	}

	viper.SetConfigFile(filepath)
	if err = viper.ReadInConfig(); err != nil {
		return
	}

	return nil
}
