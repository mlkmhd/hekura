package pkg

import (
	"fmt"
	"os"
)

func WriteToFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
