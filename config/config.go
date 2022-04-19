package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string `json:"db_username" env:"MYSQL_USER"`
	DB_PASSWORD string `json:"db_password" env:"MYSQL_PASSWORD"`
	DB_HOST     string `json:"db_host" env:"MYSQL_HOST"`
	DB_PORT     string `json:"db_port" env:"MYSQL_PORT"`
	DB_NAME     string `json:"db_name" env:"MYSQL_DBNAME"`
}

func GetConfig() Configuration {
	conf := Configuration{}

	gonfig.GetConf("config/config.json", &conf)
	return conf
}
