package main

import (
	"../lib/logger"
)



func main() {
	print("Link Start!")

	logger.SetupLogger(&logger.Settings{
		Path:	"logs",
		Name:	"goutapi",
		Ext:	"log",
		TimeFormat: "2006-01-02",
	})
	logger.Error("hahah")
}