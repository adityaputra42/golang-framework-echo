package config

import (
	"golang_framework_echo/helper"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	Db_Name     string
}

func GetConfiguration() Configuration {
	conf := Configuration{}
	err := gonfig.GetConf("config/config.json", &conf)
	helper.PanicIfError(err)
	return conf
}
