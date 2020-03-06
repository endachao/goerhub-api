package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// init config
	err := initConfig()
	if err != nil {
		panic(err)
	}
	// init App
	r := initRouter()
	_ = r.Run(fmt.Sprintf("%s:%s", viper.GetString("app.host"), viper.GetString("app.port")))
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/goerhub")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}
