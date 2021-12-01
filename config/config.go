package config

import (
	"encoding/json"
	"fmt"
	cmd "fs/cmdfunctions"
	"fs/util"
	"os"

	"github.com/fatih/color"
)

var UserHomeDir, _ = os.UserHomeDir()

//Root config
type RootConfigStruct struct {
	GlobalBin    string
	GlobalConfig string
	UserBin      string
	UserConfig   string
}

var RootConfigDefault = RootConfigStruct{
	GlobalBin:    "/usr/share/fennecshell/bin",
	GlobalConfig: "/usr/share/fennecshell/config.json",
	UserBin:      UserHomeDir + "/.local/share/fennecshell/bin",
	UserConfig:   UserHomeDir + "/.local/share/fennecshell/config.json",
}

var RootConfigDir = "/usr/config/fennecshell/"
var RootConfigFile = "config.json"
var RootConfigPath = RootConfigDir + RootConfigFile
var RootConfig RootConfigStruct
var fRoot,_ = util.GetValue(RootConfigPath);
var _ = json.Unmarshal(fRoot, &RootConfig)

//Global Config
type GlobalConfigStruct struct {
	PS1    string
	Prefix string
}

var GlobalConfigDefault = GlobalConfigStruct{
	PS1:    "%user%> ",
	Prefix: "/",
}

var GlobalConfigDir = "/usr/share/fennecshell/"
var GlobalConfigFile = "config.json"
var GlobalConfigPath = GlobalConfigDir + GlobalConfigFile
var GlobalConfig GlobalConfigStruct
var fGlobal,_ = util.GetValue(GlobalConfigPath)
var _ = json.Unmarshal(fGlobal, &GlobalConfig)

//User config
type UserConfigStruct struct {
	PS1    string
	Prefix string
}

var UserConfigDir = UserHomeDir + "/.local/share/fennecshell/"
var UserConfigFile = "config.json"
var UserConfigPath = UserConfigDir + UserConfigFile
var UserConfig UserConfigStruct
var fUser,_ = util.GetValue(UserConfigPath)
var _ = json.Unmarshal(fUser, &UserConfig)

func Run() {
	rootDefaultFile, _ := json.Marshal(RootConfigDefault)
	globalDefaultFile, _ := json.Marshal(GlobalConfigDefault)
	userDefaultFile := globalDefaultFile

	if p,_ := util.Exist(RootConfigPath); !p {
		magenta := color.New(color.Bold, color.BgHiMagenta).SprintFunc()
		fmt.Println(magenta("First run! Welcome to Fennec Shell"), "\nWe need root privileges for setup files")

		cmd.Mkdir(RootConfigDir, true)
		cmd.Touch(RootConfigPath, true)
		cmd.Echo(string(rootDefaultFile), RootConfigPath, ">", true)
		fRoot,_ = util.GetValue(RootConfigPath)
		json.Unmarshal(fRoot, &RootConfig)
	}
	if p,_ := util.Exist(GlobalConfigPath); !p {
		cmd.Mkdir(GlobalConfigDir, true)
		cmd.Touch(GlobalConfigPath, true)
		cmd.Echo(string(globalDefaultFile), RootConfig.GlobalConfig, ">", true)
	}
	if p,_ := util.Exist(UserConfigPath); !p {
		cmd.Mkdir(UserConfigDir, false)
		cmd.Touch(UserConfigPath, false)
		cmd.Echo(string(userDefaultFile), RootConfig.UserConfig, ">", false)
	}
}
