package boot

import (
	"go-web-template/conf"
	"go-web-template/pkg/database"
	"go-web-template/pkg/logger"
	"go-web-template/pkg/viperlib"
)

func Initialize() {
	viperlib.InitConfig()
	conf.LoadConf()
	logger.InitLogger(conf.LoggerFileName,
		conf.LoggerMaxSize,
		conf.LoggerMaxBackup,
		conf.LoggerMaxAge,
		conf.LoggerCompress,
		conf.LoggerLogType,
		conf.LoggerLevel)

	database.InitDB()

	database.DB.AutoMigrate()

}
