package conf

import "go-web-template/pkg/viperlib"

func LoadConf() {
	DBHost = viperlib.GetString("database.host")
	DBPort = viperlib.GetString("database.port")
	DBName = viperlib.GetString("database.dbname")
	DBUsername = viperlib.GetString("database.username")
	DBPassword = viperlib.GetString("database.password")
	DBCharset = viperlib.GetString("database.charset")
	DBType = viperlib.GetString("database.type")

	LoggerFileName = viperlib.GetString("log.filename")
	LoggerMaxSize = viperlib.GetInt("log.maxsize")
	LoggerMaxBackup = viperlib.GetInt("log.maxBackUp")
	LoggerMaxAge = viperlib.GetInt("log.maxAge")
	LoggerCompress = viperlib.GetBool("log.compress")
	LoggerLogType = viperlib.GetString("log.logType")
	LoggerLevel = viperlib.GetString("log.level")

	ServerPort = viperlib.GetString("server.port")
}
