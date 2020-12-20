package main

import (
	"fmt"
)

func main() {
	fmt.Println("csh api")
}

type Script struct {
	name      string
	script    string
	userName  string
	userEmail string
}
