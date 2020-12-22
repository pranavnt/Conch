package main

import (
	"fmt"
	"log"
	"github.com/urfave/cli/v2"
	"os"
)
func main() {
	fmt.Println("conch cli")
	app:= &cli.App{
		Name:"hello",
		Usage:"says it back",
		Action: func(c * cli.Context) error {
			fmt.Println("hey")
			return nil
		},
	}

	err:=app.Run(os.Args)
		if err!= nil {
			log.Fatal(err)
		}
}
