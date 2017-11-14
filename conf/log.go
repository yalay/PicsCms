package conf

import (
	"os"
	"path/filepath"

	"github.com/op/go-logging"
)

var Log *logging.Logger

func InitLog(logPath string, debug bool) {
	logName := filepath.Base(logPath)
	Log = logging.MustGetLogger(logName)

	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("open log file error:" + err.Error())
	}

	logBackend := logging.NewBackendFormatter(
		logging.NewLogBackend(logFile, "", 0),
		logging.MustStringFormatter(
			`[%{time:2006-01-02 15:04:05}] [%{level:.4s}] %{shortfile}:%{message}`,
		))
	logging.SetBackend(logBackend)

	if debug {
		logging.SetLevel(logging.DEBUG, logName)
	} else {
		logging.SetLevel(logging.INFO, logName)
	}
}
