package handler

import (
	"encoding/json"
	"fmt"
	cmd "fs/cmdfunctions"
	"fs/commands"
	"fs/setup"
	"fs/util"
	"io/ioutil"
	"os"
	"regexp"

	//"regexp"
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
	var args []string
	if strings.HasPrefix(command, "/$") {
		args = strings.Split(command[2:], " ")
		if command != "/$" {
			bashCmd := args[0]
			bashArgs := args[1:]
			fmt.Println(strings.Join(bashArgs, " "))
			cmd, err := cmd.RunCommand(bashCmd+" $arg", false, "", strings.Join(bashArgs, " "))
			util.Error(err)
			cmd.Run()
		} else {
			err := fmt.Errorf("Give bash command.")
			util.Error(err)
		}
	} else if strings.HasPrefix(command, "/") {
		re := regexp.MustCompile(`"(.+?)"|'(.+?)'|(\S+)`)
		var args []string
		res := re.FindAllStringSubmatch(command[1:], -1)
		for _, list := range res {
			list := list[1:]
			args = append(args, strings.Join(list, ""))
		}
		if len(args) >= 1 {
			cmd := args[0]
			cmdArgs := args[1:]
			if commands.Cmds(cmd, strings.Join(cmdArgs, " ")) {

			} else if p, _ := util.Exist(setup.RootConfig.GlobalBin + "/" + cmd); p {
				cmdDir := setup.RootConfig.GlobalBin + "/" + cmd
				cmdConfig := cmdDir + "/" + setup.GlobalConfigFile
				if p, _ := util.Exist(cmdConfig); !p {
					err := fmt.Errorf("Config doesn't exist.")
					util.Error(err)
				} else {
					cmdFile, err := util.GetValue(cmdConfig)
					var cmdJson Cmd
					util.Error(err)
					json.Unmarshal(cmdFile, &cmdJson)
					fmt.Println(cmdJson)
				}
			} else if p, _ := util.Exist(setup.RootConfig.UserBin + "/" + cmd); p {
				cmdDir := setup.RootConfig.UserBin + "/" + cmd
				cmdConfig := cmdDir + "/" + setup.UserConfigFile
				if p, _ := util.Exist(cmdConfig); !p {
					err := fmt.Errorf("Config doesn't exist.")
					util.Error(err)
				} else {
					cmdFile, err := util.GetValue(cmdConfig)
					var cmdJson Cmd
					util.Error(err)
					json.Unmarshal(cmdFile, &cmdJson)
					fmt.Println(cmdJson)
				}
			} else {
				if !commands.Cmds(cmd, strings.Join(cmdArgs, " ")) {
					err := fmt.Errorf("Command '%v' doesn't exist.", cmd)
					util.Error(err)
				}
			}
		} else {
			err := fmt.Errorf("Give command.")
			util.Error(err)
		}
	} else {
		fmt.Println(command)
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
