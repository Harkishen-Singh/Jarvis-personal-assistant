package utils

import (
	"github.com/op/go-logging"
)
var (
	log = logging.MustGetLogger("Jarvis Logger")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
)

func xfirst() {
	logging.SetFormatter(format)
}

// LoggerInfo Logging ideal information
func LoggerInfo(message string) {
	xfirst()
	log.Info(message)
}

// LoggerWarn Logging ideal warnings
func LoggerWarn(message string) {
	xfirst()
	log.Warning(message)
}

// LoggerErr Logging ideal errors
func LoggerErr(message string) {
	xfirst()
	log.Error(message)
}

// LoggerCritic Logging ideal critical information
func LoggerCritic(message string) {
	xfirst()
	log.Critical(message)
}