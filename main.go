package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printMenu()
		return
	}

	command := args[0]

	switch command {
	case "create":
		create(args[1:])
	case "help":
		help(args[1:])
	default:
		printUnknownCommand(command)
	}
}

func printMenu() {
	fmt.Println(`Goliath is a tool for help you to develop go project as a clean architecture

Usage:
	goliath <command> [arguments]

The <command> are:
	
	create		create something 
	version 	print Goliath version

Use "goliath help <command>" for more information about a command.
		`)
}

func printUnknownCommand(command string) {
	s := fmt.Sprintf(`goliath %s: unknown command
Run 'goliath help' for usage.`, command)
	fmt.Println(s)
}

func help(args []string) {
	if len(args) == 0 {
		printMenu()
	}
}
