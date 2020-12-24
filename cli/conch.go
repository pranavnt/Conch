package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("You have to provide a command :)")
		os.Exit(1)
	} else if len(os.Args) == 1 {
		fmt.Println("Welcome to Conch!")
		// add basic usage instructions
		os.Exit(1)
	}

	switch os.Args[1] {
	case "upload":
		if len(os.Args) == 4 {
			uploadScript(os.Args[2], os.Args[3])
		} else {
			fmt.Println("You must format this as `conch upload filePath scriptName` ")
		}

	case "run":
		if len(os.Args) == 3 {
			runScript(os.Args[2])
		} else {
			fmt.Println("You must format this as `conch run scriptName`")
		}

	default:
		fmt.Println("command not recognized")
		os.Exit(1)
	}
}

func uploadScript(filePath string, name string) {
	// convert file of script to string
	file, err := ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	scriptString := ""

	lines := strings.Split(file, "\n")

	for i := 0; i < len(lines)-1; i++ {
		scriptString += lines[i] + " && "
	}

	scriptString += lines[len(lines)-1]

	fmt.Println(scriptString)

	// now post scriptstring and name to the API

	// print out the api response
}

func constructURL(name string, cmd string) (url string) {
	cmd = strings.ReplaceAll(cmd, " ", "%20")

	url = "http://127.0.0.1:3000/upload?name=" + name + "&script=" + cmd

	return
}

func runScript(name string) {
	// call API to get script and run it

	resp, err := http.Get("http://127.0.0.1:3000/scripts/" + name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	data := Script{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(err)
	}

	runCmd(data.Script)
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
			log.Fatal(err)
		}
	}
	return
}

type Script struct {
	Script string `json:"script"`
}

// IO Utils
func ReadFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
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
