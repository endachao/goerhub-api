package config

import "github.com/spf13/viper"

func InitConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	return viper.ReadInConfig()
}
