package utils

import (
	"fmt"
	"os"
)

func ReadFileFromText(filename string) (string, error) {

	content, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}
	return string(content), nil

}

// Write a function save text on Text File
func WriteToFile(filename, text string) error {
	// Write the text to a file
	err := os.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	filename := "output.txt"
	text := "Hello, world!"

	err := WriteToFile(filename, text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Text written to file successfully!")
}
