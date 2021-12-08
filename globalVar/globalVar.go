package globalVar

import (
	"os"
	"os/user"
	"strings"
)

var User, _ = user.Current()

func GetDir() string {
	dir,_ := os.Getwd()
	dirHome,_ := os.UserHomeDir()
	return strings.ReplaceAll(dir, dirHome, "~")
}
