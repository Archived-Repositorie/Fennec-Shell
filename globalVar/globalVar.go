package globalVar

import (
	"os"
	"strings"
)
func GetDir() string {
	dir,_ := os.Getwd()
	dirHome,_ := os.UserHomeDir()
	return strings.ReplaceAll(dir, dirHome, "~")
}