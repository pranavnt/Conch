package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	runCmd("echo hi && echo bye")
}

func uploadScript(fileName string, name string) {

}

func runScript(name string) {
	// call API to get script

	// send GET request to server -- serverurl/scripts/name
	url := "serverURL.com/scripts/" + name
	fmt.Println(url)

	// once you have the script, do runCmd to run the script
}

func runCmd(cmd string) {
	// splits string of script into lines
	lines := strings.Split(cmd, " && ")

	// iterates through lines
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")

		command := exec.Command(line[0], line[1:]...)

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()

		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}
