package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey string
	ApiSecret string
	LogFile string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to load file %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey: cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_key").String(),
		LogFile: cfg.Section("gotrading").Key("log_file").String(),
	}
}