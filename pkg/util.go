package pkg

import (
	"os"
)

func WriteToFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		Logger.Fatalf("Error creating file with name %v: %v", fileName, err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		Logger.Fatalf("Error writing content to file %v: %v", fileName, err)
		return
	}
}