package handler

import (
	"encoding/json"
	"fmt"
	cmd "fs/cmdfunctions"
	"fs/commands"
	"fs/globalVar"
	"fs/setup"
	"fs/util"
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

func Bash(args []string, command string) {
	args = strings.Split(command[6:], " ")
	if command != "/bash " {
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
}

func SplitCmd(args *[]string, command string) {
	re := regexp.MustCompile(`"(.+?)"|'(.+?)'|(\S+)`)
	res := re.FindAllStringSubmatch(command[1:], -1)
	for _, list := range res {
		list := list[1:]
		*args = append(*args, strings.Join(list, ""))
	}
}

func GlobalBin(cmd string) {
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
}

func UserBin(cmd string) {
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
}

func Terminal(command string) {
	var args []string
	if strings.HasPrefix(command, "/bash ") {
		Bash(args, command)
	} else if strings.HasPrefix(command, "/") {
		SplitCmd(&args, command)
		if len(args) >= 1 {
			cmd := args[0]
			cmdArgs := args[1:]
			if commands.Cmds(cmd, strings.Join(cmdArgs, " ")) {
			} else if p, _ := util.Exist(setup.RootConfig.GlobalBin + "/" + cmd); p {
				GlobalBin(cmd)
			} else if p, _ := util.Exist(setup.RootConfig.UserBin + "/" + cmd); p {
				UserBin(cmd)
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
		fmt.Printf("%v: %v\n", globalVar.User.Username, command)
	}
}
