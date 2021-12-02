package cmdfunctions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Response(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}

func RunCommand(command string, root bool, custom string,arg ...string) (*exec.Cmd, error) {
	var sudo string = "bash -c"
	if root {
		sudo = "sudo bash -c"
	}
	cmds := fmt.Sprintf("%v arg='%v'; %v \"%v\"", custom,strings.Join(arg, " "), sudo, command)
	cmd := exec.Command("bash", "-c", cmds)
	Response(cmd)
	err := cmd.Run()
	return cmd, err
}

func Mkdir(dir string, root bool) error {
	command := "mkdir -p $arg"
	cmd, err := RunCommand(command, root, "",dir)

	Response(cmd)
	return err
}

func Touch(pathToFile string, root bool) error {
	command := "touch $arg"
	cmd, err := RunCommand(command, root, "",pathToFile)

	Response(cmd)
	return err
}

func Echo(input string, output string, typeChange string, root bool) error {
	command := "echo -e '$cmd' $arg"
	cmd, err := RunCommand(command, root, fmt.Sprintf("cmd='%v';", input), typeChange, output)

	Response(cmd)
	return err
}
