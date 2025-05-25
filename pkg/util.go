package pkg

import (
	"fmt"
	"os"
)

// WriteToFile creates a file with the given fileName and writes the provided content to it.
// It returns an error if file creation or writing fails.
func WriteToFile(fileName string, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", fileName, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing content to file %s: %w", fileName, err)
	}
	return nil
}
