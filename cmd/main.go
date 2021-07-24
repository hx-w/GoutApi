package main

import (
	"goutapi/pkg/logger"
)

func main() {
	logger.SetupLogger(&logger.Settings{
		Path:       "logs",
		Name:       "GoutApi",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})
	logger.Debug("logger setup success!")
}
