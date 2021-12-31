package main

import (
	"GoBitcoinTrader/config"
	"GoBitcoinTrader/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}