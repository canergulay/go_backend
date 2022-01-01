package customlogger

import (
	"log"

	"github.com/natefinch/lumberjack"
)

var (
	fileName   string = "/var/log/myapp/logs.log"
	maxSize    int    = 500
	maxBackUps int    = 3
	maxAge     int    = 28
	compress   bool   = true
)

type CustomLogger *log.Logger

func New() CustomLogger {
	logger := log.Default()
	// We will use Lumberjack to utilise our Logger
	logger.SetOutput(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackUps,
		MaxAge:     maxAge,
		Compress:   compress,
	})
	return logger
}
