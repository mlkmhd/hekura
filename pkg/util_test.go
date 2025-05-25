package pkg

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteToFile_Success(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_output.txt")
	testContent := "Hello, Hekura!"

	err := WriteToFile(testFile, testContent)
	if err != nil {
		t.Fatalf("WriteToFile() returned an unexpected error: %v", err)
	}

	content, readErr := os.ReadFile(testFile)
	if readErr != nil {
		t.Fatalf("Failed to read back test file: %v", readErr)
	}

	if string(content) != testContent {
		t.Errorf("WriteToFile() content = %q, want %q", string(content), testContent)
	}
}

func TestWriteToFile_ErrorPathIsDirectory(t *testing.T) {
	tmpDir := t.TempDir() // This directory path will be used as the "file" to write to
	testContent := "This should not be written."

	err := WriteToFile(tmpDir, testContent)
	if err == nil {
		t.Fatalf("WriteToFile() did not return an error when trying to write to a directory")
	}

	// Check if the error message indicates a creation or writing issue.
	// The exact error message might vary by OS ("is a directory", "permission denied", etc.)
	// So, we check for common substrings.
	errMsg := err.Error()
	if !strings.Contains(errMsg, "error creating file") && !strings.Contains(errMsg, "error writing content") {
		t.Errorf("WriteToFile() error message = %q, want to contain 'error creating file' or 'error writing content'", errMsg)
	}

	// Ensure no file was created with the directory's name (though it shouldn't be possible)
	// and that the directory itself still exists and is a directory.
	info, statErr := os.Stat(tmpDir)
	if statErr != nil {
		t.Fatalf("os.Stat() on tmpDir failed: %v", statErr)
	}
	if !info.IsDir() {
		t.Errorf("tmpDir is no longer a directory after WriteToFile call")
	}
}

// Note: The InitLoggerForTest and NewLogger helper functions and osExit mocking
// are no longer needed here as WriteToFile now returns an error directly.
// They were removed as part of this overwrite.
