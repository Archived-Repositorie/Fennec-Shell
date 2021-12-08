//IM VERY ANGY AT THIS NOW

package main

import (
	//cmd "fs/cmdfunctions"
	"flag"
	"fmt"
	"fs/globalVar"
	"fs/handler"
	"fs/setup"
	"fs/util"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	script := flag.Bool("c", false, "run command inside string")
	flag.Parse()
	if *script {
		cmdArgs := flag.Args()
		if len(cmdArgs) > 0 {
			handler.Terminal(strings.Join(cmdArgs, " "))
		} else {
			fmt.Println(cmdArgs)
		}
	} else {
		command()
	}
}

type PS1Tags struct {
	Tag     string
	Replace string
}

func command() {
	os.Chdir(globalVar.User.HomeDir)
	setup.Run()
	textBold := color.New(color.Bold)
	textBold.Println("Welcome to Fennec Shell!")

	var input string
	dir := globalVar.GetDir()
	for {
		PS1 := setup.UserConfig.PS1
		PS1Replace := []PS1Tags{
			{
				Tag:     "%user%",
				Replace: globalVar.User.Username,
			},
			{
				Tag:     "%dir%",
				Replace: dir,
			},
		}
		for _, PS1Tag := range PS1Replace {
			PS1 = strings.ReplaceAll(PS1, PS1Tag.Tag, PS1Tag.Replace)
		}
		fmt.Print(PS1)
		util.Scanner(&input)
		handler.Terminal(input)
		dir = globalVar.GetDir()
	}
}
