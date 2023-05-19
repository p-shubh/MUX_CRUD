package db

import (
	"CEO_ASSIST/config"
	"flag"
)

var configFile = flag.String("f", "config/DataCollector.json", "path for the configfile")

func DATABASE_CONNECTION() *config.DataCollector {

	var DataConfig = config.NewConfig(*configFile)
	return DataConfig
}
