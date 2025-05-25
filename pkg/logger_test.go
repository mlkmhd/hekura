package pkg

import (
	"bytes"
	// "os" // No longer needed for TestMain or direct verbose checks here
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

// ensureLoggerInitialized initializes the global Logger if it's nil.
// It sets the default output to a bytes.Buffer to discard logs during tests.
func ensureLoggerInitialized() {
	if Logger == nil {
		Logger = logrus.New()
		Logger.SetLevel(logrus.InfoLevel) // Default level
	}
	// For tests, ALWAYS set output to a discard buffer by default,
	// overriding any os.Stdout set by logger.go's init().
	// Tests needing to capture output can then redirect this.
	Logger.SetOutput(&bytes.Buffer{}) 
}

func TestSetLogLevel_ValidLevels(t *testing.T) {
	ensureLoggerInitialized()

	testCases := []struct {
		name          string
		levelStr      string
		expectedLevel logrus.Level
	}{
		{"debug", "debug", logrus.DebugLevel},
		{"info", "info", logrus.InfoLevel},
		{"warn", "warn", logrus.WarnLevel},
		{"error", "error", logrus.ErrorLevel},
		{"fatal", "fatal", logrus.FatalLevel},
		{"panic", "panic", logrus.PanicLevel},
	}

	// Save and restore original logger settings around all test cases in this function
	originalOutput := Logger.Out
	originalLevel := Logger.Level
	defer func() {
		Logger.SetOutput(originalOutput)
		Logger.SetLevel(originalLevel)
	}()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Each t.Run can further isolate if needed, but for SetLogLevel,
			// modifying the global Logger's level is what we're testing.
			SetLogLevel(tc.levelStr)
			if Logger.GetLevel() != tc.expectedLevel {
				t.Errorf("SetLogLevel(%q) = %v, want %v", tc.levelStr, Logger.GetLevel(), tc.expectedLevel)
			}
		})
	}
}

func TestSetLogLevel_InvalidLevel(t *testing.T) {
	ensureLoggerInitialized() // Ensures global Logger is not nil, but we'll override it

	// Store the original global Logger instance
	originalGlobalLogger := Logger
	
	// Create a new logger instance specifically for this test
	testLogger := logrus.New()
	var logCapture bytes.Buffer
	testLogger.SetOutput(&logCapture)
	// Set level to WarnLevel or DebugLevel to ensure Warnf messages are captured.
	// The function SetLogLevel will later change it to InfoLevel for the invalid case.
	testLogger.SetLevel(logrus.WarnLevel) 

	// Replace the global Logger with our testLogger
	Logger = testLogger
	
	// Defer restoration of the original global Logger
	defer func() {
		Logger = originalGlobalLogger
	}()

	// Set a known level (on testLogger, which is now global Logger)
	// This line is redundant if testLogger.SetLevel above is sufficient.
	// Logger.SetLevel(logrus.ErrorLevel) 

	invalidLevelStr := "ThisIsNotALevel"
	SetLogLevel(invalidLevelStr)

	if Logger.GetLevel() != logrus.InfoLevel {
		t.Errorf("SetLogLevel(%q) did not default to InfoLevel, got %v", invalidLevelStr, Logger.GetLevel())
	}

	logStr := logCapture.String()
	expectedWarningMsg := "Unknown log level specified ('ThisIsNotALevel'), defaulting to 'info'"
	if !strings.Contains(logStr, expectedWarningMsg) {
		t.Errorf("SetLogLevel(%q) did not log the expected warning.\nExpected to contain: %q\nGot log: %s", invalidLevelStr, expectedWarningMsg, logStr)
	}
}

func TestSetLogLevel_CaseInsensitive(t *testing.T) {
	ensureLoggerInitialized()

	originalOutput := Logger.Out
	originalLevel := Logger.Level
	defer func() {
		Logger.SetOutput(originalOutput)
		Logger.SetLevel(originalLevel)
	}()
	
	// Discard logs for these test cases as we only check the level, not output
	Logger.SetOutput(&bytes.Buffer{})


	testCases := []struct {
		name          string
		levelStr      string
		expectedLevel logrus.Level
	}{
		{"DEBUG", "DEBUG", logrus.DebugLevel},
		{"Info", "Info", logrus.InfoLevel},
		{"wArN", "wArN", logrus.WarnLevel},
		{"ERROR", "ERROR", logrus.ErrorLevel},
		{"FaTaL", "FaTaL", logrus.FatalLevel},
		{"pAnIc", "pAnIc", logrus.PanicLevel},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SetLogLevel(tc.levelStr)
			if Logger.GetLevel() != tc.expectedLevel {
				t.Errorf("SetLogLevel(%q) = %v, want %v", tc.levelStr, Logger.GetLevel(), tc.expectedLevel)
			}
		})
	}
}
