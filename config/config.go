package config

import (
	"encoding/json"
	"fmt"
	cmd "fs/cmdfunctions"
	"fs/util"

	"github.com/fatih/color"
)

type ConfigStruct struct {
	GlobalBin string
	GlobalConfig string
	UserBin string
	UserConfig string
}

var ConfigDefault = ConfigStruct{
	GlobalBin: "/usr/share/fennecshell/bin",
	GlobalConfig: "/usr/share/fennecshell/config.json",
	UserBin: "~/.local/share/fennecshell/bin",
	UserConfig: "~/.local/share/fennecshell/config.json",
}

var ConfigDir = "/usr/config/fennecshell/"
var ConfigFile = "config.json"
var ConfigPath = ConfigDir+ConfigFile
var Config = ConfigDefault

var ConfigJson = json.Unmarshal(util.GetValue(ConfigDir), &Config)



func Run() {
	if !util.Exist(ConfigPath) {
		magenta := color.New(color.Bold, color.BgHiMagenta).SprintFunc()
		output, err := json.Marshal(ConfigDefault)

		fmt.Println(magenta("First run! Welcome to Fennec Shell"), "\nWe need root privileges for setup files")
		cmd.Mkdir(ConfigDir, true)

		util.Error(err)

		cmd.Touch(ConfigPath, true)
		cmd.Echo(string(output), ConfigPath, ">", true)
	}
}