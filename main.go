//IM VERY ANGY AT THIS NOW

package main

import (
	//cmd "fs/cmdfunctions"
	"fmt"
	"fs/config"
	"fs/handler"
	"fs/util"
	"os/user"
	"strings"

	"github.com/fatih/color"
)

type PS1Tags struct {
	Tag string
	Replace string
}
var User,_ = user.Current()

var PS1Replace = []PS1Tags{
	PS1Tags{
		Tag: "%user",
		Replace: User.Name,
	},
}

func main() {
	config.Run()
	textBold := color.New(color.Bold)
	textBold.Println("Welcome to Fennec Shell!")
	PS1 := config.UserConfig.PS1
	fmt.Printf("\n"+PS1)
	for _,PS1Tag := range PS1Replace {
		PS1 = strings.ReplaceAll(PS1, PS1Tag.Tag, PS1Tag.Replace)
	}

	var input string
	for {
		util.Scanner(&input)
		handler.Terminal(input)
	}
	
}
