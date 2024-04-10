package conf

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Println("配置文件读取错误，请检查文件路径:", err)
		panic(err)
	}
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		log.Println(err) // 日志内容
		panic(err)
	}
	LoadServer(file)
	LoadMysqlData(file)
	LoadRedis(file)
}

func LoadServer(file *ini.File) {
	// 获取文件中的 "AppMode" 配置项的值，并将其赋给 AppMode 变量
	AppMode = file.Section("service").Key("AppMode").String()

	// 获取文件中的 "HttpPort" 配置项的值，并将其赋给 HttpPort 变量
	HttpPort = file.Section("service").Key("HttpPort").String()
}

// LoadMysqlData 函数从INI配置文件中加载MySQL数据库的相关配置信息
func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
