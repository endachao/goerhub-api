package main

import (
	"fmt"
	"github.com/spf13/viper"
	"goerhubApi/config"
	"goerhubApi/dao"
	"goerhubApi/routes"
)

func main() {
	// init config
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	// init mysql connect
	mysql := dao.MysqlConnect{}
	if err := mysql.Connect(); err != nil {
		panic(err)
	}
	defer mysql.Close()

	// init App
	r := routes.InitRouter()
	_ = r.Run(fmt.Sprintf("%s:%s", viper.GetString("app.host"), viper.GetString("app.port")))
}
