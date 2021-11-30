package cmdfunctions

import (
	"fmt"
	"fs/util"
	"os"
	"os/exec"
)

func Response(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}

func Mkdir(dir string, root bool) {
	var sudo string = ""
	if root {
		sudo = "sudo "
	}

	command := fmt.Sprintf("%vmkdir -p %v", sudo, dir)
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	util.Error(err)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}

func Touch(pathToFile string, root bool) {
	var sudo string = ""
	if root {
		sudo = "sudo "
	}

	command := fmt.Sprintf("%vtouch %v", sudo, pathToFile)
	fmt.Println(command)
	
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	util.Error(err)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}

func Echo(input string, output string, typeChange string, root bool) {
	var sudo string = ""
	if root {
		sudo = "sudo "
	}

	command := fmt.Sprintf("%vecho -e '%v' %v %v %v", sudo, input, typeChange, sudo ,output)
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	util.Error(err)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}

func Exist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	util.Error(err)
	return false
}
