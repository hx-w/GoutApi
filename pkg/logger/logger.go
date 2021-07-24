package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type Settings struct {
	Path       string `yaml:"path"`
	Name       string `yaml:"name"`
	Ext        string `yaml:"ext"` // extend
	TimeFormat string `yaml:"time-format"`
}

var (
	logFile            *os.File
	defaultPrefix      = ""
	defaultCallerDepth = 2
	logger             *log.Logger // true logger
	mutex              sync.Mutex
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
)

// redefine type
type logLevel int

const (
	DEBUG logLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

const flags = log.LstdFlags

// --------- func ---------

func initLogger() {
	logger = log.New(os.Stdout, defaultPrefix, flags)
}

func SetupLogger(settings *Settings) {
	dir := settings.Path
	fileName := fmt.Sprintf("%s-%s-%s", settings.Name, time.Now().Format(settings.TimeFormat), settings.Ext)

	logFile, err := mustOpen(fileName, dir)
	if err != nil {
		log.Fatalf("logging.Setup fatal: %s", err)
	}

	multw := io.MultiWriter(os.Stdout, logFile)
	logger = log.New(multw, defaultPrefix, flags)
}

func setPrefix(level logLevel) {
	_, file, line, ok := runtime.Caller(defaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d] ", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s] ", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

// Debug prints debug log
func Debug(v ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	setPrefix(DEBUG)
	logger.Println(v...)
}

// Info prints normal log
func Info(v ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	setPrefix(INFO)
	logger.Println(v...)
}

// Warn prints warning log
func Warn(v ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	setPrefix(WARNING)
	logger.Println(v...)
}

// Error prints error log
func Error(v ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	setPrefix(ERROR)
	logger.Println(v...)
}

// Fatal prints error log then stop the program
func Fatal(v ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	setPrefix(FATAL)
	logger.Fatalln(v...)
}
