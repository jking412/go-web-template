package database

import (
	"fmt"
	"go-web-template/conf"
	"go-web-template/pkg/logger"
	"go-web-template/pkg/viperlib"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	switch conf.DBType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			viperlib.GetString("database.username"),
			viperlib.GetString("database.password"),
			viperlib.GetString("database.host"),
			viperlib.GetString("database.port"),
			viperlib.GetString("database.dbname"),
			viperlib.GetString("database.charset"),
		)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.ErrorString("database", "init", err.Error())
			return err
		}
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(viperlib.GetString(conf.DBName)), &gorm.Config{})
		if err != nil {
			logger.ErrorString("database", "init", err.Error())
			return err
		}
	}
	logger.InfoString("database", "init", "database connected")
	return nil
}
