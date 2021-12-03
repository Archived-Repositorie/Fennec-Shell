//IM VERY ANGY AT THIS NOW

package main

import (
	//cmd "fs/cmdfunctions"
	"fmt"
	"fs/globalVar"
	"fs/handler"
	"fs/setup"
	"fs/util"
	"os"
	"os/user"
	"strings"

	"github.com/fatih/color"
)

type PS1Tags struct {
	Tag     string
	Replace string
}

var User, _ = user.Current()

func main() {
	os.Chdir(User.HomeDir)
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
				Replace: User.Name,
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
