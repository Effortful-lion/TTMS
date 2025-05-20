package config

import (
	"log"
	"github.com/spf13/viper"
)

var Conf = &Config{}

type Config struct {
	AppConfig *AppConfig `mapstructure:"app"`
	MysqlConfig  *MysqlConfig  `mapstructure:"mysql"`
	RedisConfig *RedisConfig `mapstructure:"redis"`
	JwtConfig *JwtConfig `mapstructure:"jwt"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	HttpPort int `mapstructure:"http"`
	HttpsPort int `mapstructure:"https"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

type JwtConfig struct {
	Expire int `mapstructure:"expire"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db int `mapstructure:"db"`
}

func InitConfig() (err error){
	viper.SetConfigFile("./config/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}