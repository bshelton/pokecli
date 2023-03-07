/*
Package logger creates global logger
*/
package logger

import (
	"log"
	"os"

	"github.com/bshelton/pokecli/pkg/config"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

// logLevel controls the global log level used by the logger.
var level = LevelInfo

func Level() int {
	return level
}

var Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
var FileLogger *log.Logger

func Setup(prefix string, logLevel int) {
	level = logLevel
	Logger = log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	file, err := os.OpenFile(config.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		FileLogger = log.New(file, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func Debug(v ...interface{}) {
	logFile(v...)
	if level <= LevelDebug {
		Logger.Printf("[DEBUG] %v\n", v...)
	}
}

func Info(v ...interface{}) {
	logFile(v...)
	if level <= LevelInfo {
		Logger.Printf("[INFO] %v\n", v...)
	}
}

func Warn(v ...interface{}) {
	logFile(v...)
	if level <= LevelWarning {
		Logger.Printf("[WARN] %v\n", v...)
	}
}

func Error(v ...interface{}) {
	logFile(v...)
	if level <= LevelError {
		Logger.Printf("[ERROR] %v\n", v...)
	}
}

func Fatal(v ...interface{}) {
	logFile(v...)
	if level <= LevelFatal {
		Logger.Printf("[FATAL] %v\n", v...)
		os.Exit(1)
	}
}

func logFile(v ...interface{}) {
	if FileLogger != nil {
		FileLogger.Printf("%v\n", v...)
	}
}
