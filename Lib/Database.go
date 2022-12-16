package Lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mini-wallet-exercise/Config"
)

var DB *gorm.DB

func initDB() {
	con, err := gorm.Open(postgres.Open(Config.DatabaseConnectionString()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	DB = con
}
