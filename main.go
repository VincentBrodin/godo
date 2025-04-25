package main

import (
	"fmt"
	"os"
	"path"

	"github.com/VincentBrodin/godo/pkg/engine"
	"github.com/VincentBrodin/godo/pkg/parser"
)

func main() {
	args := os.Args[1:]
	argc := len(args)

	// No args
	if argc <= 0 {
		fmt.Println("No args")
	}

	// Find godo file
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	godoPath := path.Join(dir, "godo.yml")
	data, err := os.ReadFile(godoPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	godoFile, err := parser.Parse(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	command := godoFile.Commands[args[0]]
	engine.Run(command)
}
