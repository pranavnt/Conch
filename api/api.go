package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	entry := Script{
		Name:      "pranav",
		Script:    "echo hello world!",
		UserName:  "pranavnt",
		UserEmail: "pranavnt@outlook.com",
	}

	addScript(entry)
}

func addScript(script Script) {
	scriptJSON, _ := json.Marshal(script)

	fmt.Println(string(scriptJSON))
}

type Script struct {
	Name      string
	Script    string
	UserName  string
	UserEmail string
}
