package util

import (
	"bufio"
	"os"

	"github.com/fatih/color"
)

func Error(err error) {
	if err != nil {
		c := color.New(color.BgRed, color.Bold)
		c.Printf("\nError: %v\n", err.Error())
	}
}

func Exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetValue(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	return file, err
}

func Scanner(input *string) {
	scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        *input = scanner.Text()
    }
	Error(scanner.Err())
}