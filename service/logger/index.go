package logger

import (
	"log"

	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger = logging.MustGetLogger("Jarvis Logger")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
)

func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/index.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
	logging.SetFormatter(format)
}

// Info Logging ideal information
func Info(message string) {
	logger.Info(message)
	log.Println(message)
}

// Warn Logging ideal warnings
func Warn(message string) {
	logger.Warning(message)
	log.Println(message)
}

// Error Logging ideal errors (Note: Calling this will automatically create a panic)
func Error(err error) {
	logger.Error(err)
	log.Panic(err)
}

// Critic Logging ideal critical information (Note: Calling this will stop the execution of the program)
func Critic(err error) {
	logger.Critical(err)
	log.Fatal(err)
}
