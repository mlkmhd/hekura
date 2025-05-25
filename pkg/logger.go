package pkg

import (
	"os"
	"strings" // Added import

	"github.com/sirupsen/logrus"
)

// Logger is a global logrus.Logger instance used for logging throughout the pkg package.
// It is initialized in the init function to output to os.Stdout with InfoLevel.
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

// SetLogLevel sets the logging level for the global Logger.
// It accepts a string representation of the log level (e.g., "debug", "info", "warn").
// It converts the input level to lowercase for case-insensitive comparison.
// If an unknown log level is provided, it defaults to "info" and logs a warning
// using the original (case-preserved) logLevel string.
func SetLogLevel(logLevel string) {
	lowerLogLevel := strings.ToLower(logLevel)
	switch lowerLogLevel {
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
		Logger.Warnf("Unknown log level specified ('%s'), defaulting to 'info'", logLevel)
		Logger.SetLevel(logrus.InfoLevel)
	}
}
