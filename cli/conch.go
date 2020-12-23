package main

import (
	"bytes"
	"fmt"
	"io"

	"os"
	"os/exec"
	"strings"
)

func main() {
	uploadScript("hi", "bye")
	//runCmd("echo hi && echo bye")
}

func uploadScript(filePath string, name string) {
	// convert file of script to string
	file, err := ReadFile("cli/test.sh")

	fmt.Println(file)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
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

func runScript(name string) {
	// call API to get script and run it

	// send GET request to server -- serverurl/scripts/name
	url := "serverURL.com/scripts/" + name
	fmt.Println(url)

	// once you have the script, do runCmd to run the script
}

// IO Utils
func ReadFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "nil", err
	}
	defer f.Close()

	var n int64

	if fi, err := f.Stat(); err == nil {
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	}

	return readAll(f, n+bytes.MinRead)
}

func readAll(r io.Reader, capacity int64) (b string, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))

	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return string(buf.Bytes()), err
}
