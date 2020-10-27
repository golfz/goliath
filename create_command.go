package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func create(args []string) {
	if len(args) == 0 {
		printCreateHelp()
		return
	}

	whatToCreate := args[0]

	switch whatToCreate {
	case "module":
		createModule(args[1:])
	default:
		printCreateUnknownArgument(whatToCreate)
	}

}

func printCreateHelp() {
	fmt.Println(`
Usage:
	goliath create [arguments]

The [arguments] are:

	module		create struct of a new goliath module
	`)
}

func printCreateUnknownArgument(arg string) {
	s := fmt.Sprintf(`goliath create %s: unknown command
Run 'goliath help create' for usage.`, arg)
	fmt.Println(s)
}

func createModule(args []string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	isDir, _ := isDirectory(path)
	if !isDir { // because in error case: isDir is false
		fmt.Println("it's strange, why is the current path not a directory?")
		os.Exit(1)
	}

	isDirEmpty, err := isDirectoryEmpty(path)
	if err != nil {
		fmt.Println("can not read data in current directory")
		os.Exit(1)
	}
	if !isDirEmpty {
		fmt.Println("current directory is not empty")
		os.Exit(1)
	}

	// todo: create folder structure
	dirList := []string{"data", "entity", "implement", "interface", "test", "util"}

	for _, s := range dirList {
		err = os.Mkdir(s, 0755)
		if err != nil {
			fmt.Println("can not create", s, "directory")
			os.Exit(1)
		}
	}

}

func isDirectory(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}

func isDirectoryEmpty(path string) (bool, error) {
	isDir, err := isDirectory(path)
	if err != nil {
		return false, err
	}

	if !isDir {
		return false, errors.New("this path is not a directory")
	}

	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err
}
