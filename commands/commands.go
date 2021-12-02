package commands

import (
	"fmt"
	"fs/globalVar"
	"fs/util"
	"os"
	"strings"
)

func Cmds(cmd string, args ...string) bool {
	dirRoot, _ := os.Getwd()
	switch cmd {
	case "cd":
		dir := strings.Join(args, " ")
		if p, _ := util.Exist(dirRoot + "/" + dir); !p {
			err := fmt.Errorf("Directory './%v' doesn't exit.", dir)
			util.Error(err)
		} else {
			os.Chdir(dir)
			fmt.Printf("\nChanged directory to '%v'\n", args[0])
			fmt.Printf("Now its '%v'\n", globalVar.GetDir())
		}
		return true
	default:
		return false
	}
	return false
}
