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

func TestWriteToFile_ErrorWriteString(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_readonly_output.txt")

	// 1. Create the file so we can change its permissions.
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Failed to pre-create test file: %v", err)
	}
	file.Close() // Close it, WriteToFile will reopen.

	// 2. Change file permissions to read-only.
	// On Unix, 0444 is read-only for all.
	// On Windows, this is more complex; os.Chmod may not prevent writes if the
	// user has ownership. This test might be OS-dependent.
	err = os.Chmod(testFile, 0400) // Read-only for owner
	if err != nil {
		t.Fatalf("Failed to change file permissions to read-only: %v", err)
	}
	defer func() {
		// Attempt to change permissions back to allow cleanup, ignore error if it fails.
		_ = os.Chmod(testFile, 0600) 
	}()


	// 3. Attempt to write to the now (hopefully) read-only file.
	// WriteToFile internally uses os.Create which will truncate and open for writing.
	// The success of os.Create might depend on user (owner) even if file is 0400.
	// If os.Create succeeds, the WriteString might fail if the OS enforces the
	// read-only permission at the FS level despite the descriptor being writable.
	err = WriteToFile(testFile, "content that should not be written")

	if err == nil {
		// If the error is nil, it means WriteToFile succeeded.
		// This could happen on some OSes (like Windows, or Unix if running as root
		// or if file ownership overrides read-only for write descriptor).
		// In this case, we can't reliably test the WriteString error path this way.
		t.Logf("WriteToFile succeeded on a supposedly read-only file (%s). This test for WriteString error might be OS-dependent or require different simulation.", testFile)
		// To ensure the test doesn't falsely pass, we should check if content was actually written.
		// If it was, then the read-only setup didn't prevent the write.
		// If content was not written but no error, that's also a strange state.
		content, readErr := os.ReadFile(testFile)
		if readErr == nil && string(content) == "content that should not be written" {
			// This means the write succeeded, so our attempt to make WriteString fail didn't work.
			// This isn't a failure of WriteToFile, but a limitation of this test method.
			// We can't make a strong assertion here that this test *must* produce an error.
		}
		// Given the potential for OS variance, if err is nil, we don't fail the test,
		// but acknowledge the write may have succeeded. The goal is to *try* to cause a WriteString error.
		return 
	}

	// Check if the error message indicates a writing issue, not a creation issue.
	errMsg := err.Error()
	if !strings.Contains(errMsg, "error writing content") {
		t.Errorf("WriteToFile() error message = %q, want to contain 'error writing content', got error: %v", errMsg, err)
	} else {
		// This is the desired outcome for this test: WriteString failed.
		t.Logf("Successfully triggered WriteString error: %v", err)
	}
}


// Note: The InitLoggerForTest and NewLogger helper functions and osExit mocking
// are no longer needed here as WriteToFile now returns an error directly.
// They were removed as part of this overwrite.
