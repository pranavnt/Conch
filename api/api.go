package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
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
