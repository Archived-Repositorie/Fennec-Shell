package util

import (
	"github.com/fatih/color"
	"os"
)

func Error(err error) {
	if err != nil {
		c := color.New(color.BgRed, color.Bold)
		c.Printf("\nError: %v", err.Error())
	}
}

func Exist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	Error(err)
	return false
}

func GetValue(path string) []byte {
	file, err := os.ReadFile(path)
	Error(err)
	return file
}