package viperlib

import "github.com/spf13/viper"

var Viper *viper.Viper

func InitConfig() {
	Viper = viper.New()
	// 读取项目根目录下的config.yaml文件
	Viper.SetConfigName("config")
	Viper.AddConfigPath("./config")
	Viper.AddConfigPath(".")
	Viper.SetConfigType("yaml")
	if err := Viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 自动检测配置文件变化
	Viper.WatchConfig()
	// 读取环境变量
	Viper.AutomaticEnv()
}

func GetString(key string) string {
	return Viper.GetString(key)
}

func GetStrings(key string) []string {
	return Viper.GetStringSlice(key)
}

func GetInt(key string) int {
	return Viper.GetInt(key)
}

func GetBool(key string) bool {
	return Viper.GetBool(key)
}
