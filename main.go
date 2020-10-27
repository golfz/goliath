package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0  {
		printHelp()
		return
	}

	//args[]
}

func printHelp() {
	fmt.Println(`
Goliath is a tool for help you to develop go project as a clean architecture

Usage:
	goliath <command> [arguments]

The <command> are:
	
	create		create something 
	update		update goliath project structure
	version 	print Goliath version

Use "goliath help <command>" for more information about a command.
		`)
}
