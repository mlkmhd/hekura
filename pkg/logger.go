package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	// Initialize the logger
	Logger = logrus.New()

	// Set the logger to write logs to a file
	//logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err == nil {
	//    Logger.SetOutput(logFile)
	//} else {
	//    Logger.Info("Failed to log to file, using default stderr")
	//}

	Logger.SetOutput(os.Stdout)

	// Set log level
	Logger.SetLevel(logrus.InfoLevel)
}

func SetLogLevel(logLevel string) {
	switch logLevel {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		Logger.SetLevel(logrus.FatalLevel)
	case "panic":
		Logger.SetLevel(logrus.PanicLevel)
	default:
		Logger.Warn("Unknown log level specified, defaulting to 'info'")
		Logger.SetLevel(logrus.InfoLevel)
    }
}
