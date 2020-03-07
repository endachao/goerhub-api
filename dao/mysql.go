package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"time"
)

var db *gorm.DB

type MysqlConnect struct {
}

func (m *MysqlConnect) Connect() (err error) {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"))

	db, err = gorm.Open("mysql", dbUrl)
	if err != nil {
		return err
	}
	log.Println("Connect mysql successfully")
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	return nil
}

func (m *MysqlConnect) Close() {
	_ = DB().Close()
}

func DB() *gorm.DB {
	return db
}
