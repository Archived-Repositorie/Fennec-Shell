package cmdfunctions

import (
	"fmt"
	"os"
	"os/exec"
)

func Response(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}

func RunCommand(command string) (*exec.Cmd, error) {
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	return cmd,err
}

func Mkdir(dir string, root bool) error {
	var sudo string = ""
	if root {
		sudo = "sudo "
	}

	command := fmt.Sprintf("%vmkdir -p %v", sudo, dir)
	cmd,err := RunCommand(command)

	Response(cmd)
	return err
}

func Touch(pathToFile string, root bool) error {
	var sudo string = ""
	if root {
		sudo = "sudo "
	}

	command := fmt.Sprintf("%vtouch %v", sudo, pathToFile)
	cmd,err := RunCommand(command)

	Response(cmd)
	return err
}

func Echo(input string, output string, typeChange string, root bool) error {
	var sudo string = ""
	if root {
		sudo = "sudo "
	}

	command := fmt.Sprintf("%vecho -e '%v' %v %v %v", sudo, input, typeChange, sudo ,output)
	cmd,err := RunCommand(command)

	Response(cmd)
	return err
}