package handler

import (
	"encoding/json"
	"fmt"
	cmd "fs/cmdfunctions"
	"fs/util"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Cmd struct {
	Core             string
	Command          string
	ShortDescription string
	LongDescription  string
	Arguments        string
}

func Terminal(command string) {
	var re = regexp.MustCompile(`(?m)\w+|\"\"[\w\s]*`)
	args := re.FindAllString(command, -1)
	if strings.HasPrefix(command, "/$") {
		bashCmd := args[0]
		bashArgs := append(args, bashCmd)
		cmd, err := cmd.RunCommand(bashCmd, false, strings.Join(bashArgs, " "))
		util.Error(err)
		cmd.Run()
	} else if strings.HasPrefix(command, "/") {
		fmt.Println(args)
	}
}

func Test() {
	root := "./test.bin/"
	folders, _ := ioutil.ReadDir(root)
	for _, f := range folders {
		path := root + f.Name() + "/"
		cmdFile, err := os.ReadFile(path + "cmd.json")
		util.Error(err)
		var cmdJson Cmd
		json.Unmarshal(cmdFile, &cmdJson)
		_, err = cmd.RunCommand(strings.ReplaceAll(cmdJson.Command, "$core", path+cmdJson.Core), false, "")
		util.Error(err)
	}
}
