package tests

import (
	"GoutApi/lib/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	logger.SetupLogger(&logger.Settings{
		Path:       "logs",
		Name:       "GoutApi",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})

	logger.Debug("DEBUG test")
	logger.Info("INFO test")
	logger.Warn("WARNING test")
	logger.Error("ERROR test")
	logger.Fatal("FATAL test")
}
