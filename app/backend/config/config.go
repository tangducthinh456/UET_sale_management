package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	config *viper.Viper

	serverconfig *ServerConfig
	dbconfig *DBConfig
	redisconfig *RedisConfig
)

func Init(){
	InitConfig()
	serverConfigInit()
	dbConfigInit()
	redisConfigInit()
}

func InitConfig(){
    config = viper.New()
    config.SetConfigName("config")
    config.SetConfigType("yaml")
	config.AddConfigPath(".")
    config.AddConfigPath("./config/")
	config.AddConfigPath("../config/")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

type ServerConfig struct{
	ServerHost string
	ServerPort string
}

func GetServerConfig() *ServerConfig{
	return serverconfig
}

func serverConfigInit() {
	serverconfig = &ServerConfig{
		ServerHost: config.GetString("server.Host"),
		ServerPort: config.GetString("server.Port"),
	}
}

type DBConfig struct {
	Host string
	Port string
	DBName string
	Username string
	Password string
}

func GetDBConfig() *DBConfig{
	return dbconfig
}

func dbConfigInit() {
	dbconfig = &DBConfig{
		Host:     config.GetString("database.Host"),
		Port:     config.GetString("database.Port"),
		DBName:   config.GetString("database.database"),
		Username: config.GetString("database.user"),
		Password: config.GetString("database.Password"),
	}
}

type RedisConfig struct {
	Host string
	Port string
	Password string
}

func GetRedisConfig() *RedisConfig{
	return redisconfig
}

func redisConfigInit() {
	redisconfig = &RedisConfig{
		Host:     config.GetString("reids.Host"),
		Port:     config.GetString("redis.Port"),
		Password: config.GetString("redis.Password"),
	}
}