package util

import (
	"github.com/fatih/color"
)

func Error(err error) {
	if err != nil {
		c := color.New(color.BgRed, color.Bold)
		c.Printf("\nError: %v", err.Error())
	}
}