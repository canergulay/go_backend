package logclient

import (
	log "github.com/sirupsen/logrus"

	"github.com/natefinch/lumberjack"
)

var (
	fileName   string = "/var/log/myapp/logs.log"
	maxSize    int    = 500
	maxBackUps int    = 3
	maxAge     int    = 28
	compress   bool   = true
)

var Log *log.Logger

func InitLogger() {

	logger := log.New()
	// We will use Lumberjack to utilise our Logger
	logger.SetOutput(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackUps,
		MaxAge:     maxAge,
		Compress:   compress,
	})

	Log = logger
}
