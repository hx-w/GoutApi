package main

import (
	"GoutApi/lib/logger"
)

func main() {
	logger.SetupLogger(&logger.Settings{
		Path:       "logs",
		Name:       "GoutApi",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})
	logger.Error("A TEST")
}
