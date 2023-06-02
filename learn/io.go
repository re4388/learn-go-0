package learn

import (
	"fmt"
	"os"
)

func CreateFile(fileName string) {
	file, _ := os.Create(fileName)
	fmt.Fprint(file, "qqqq")
	file.Close()
}

func AppendFile(fn, text string) error {
	file, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file for writing: %w", err)
	}

	defer file.Close()

	if _, err := file.Write([]byte(text)); err != nil {
		return fmt.Errorf("error writing text to fiel: %w", err)
	}

	return nil
}
