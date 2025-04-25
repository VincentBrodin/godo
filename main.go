package main

import (
	"fmt"
	"os"

	"github.com/VincentBrodin/godo/pkg/engine"
	"github.com/VincentBrodin/godo/pkg/parser"
	"github.com/VincentBrodin/godo/pkg/utils"
)

func main() {
	// Start of by reading the file
	data, err := utils.ReadByName("godo", ".exe", ".exe~", ".dll", ".so", ".dylib", ".test", ".out")
	if err != nil {
		fmt.Println("No godo file found!")
		return
	}

	godoFile, err := parser.Parse(data)
	if err != nil {
		fmt.Printf("Could not parse godo file: %s", err)
		return
	}

	args := os.Args[1:]
	argc := len(args)

	// No args
	if argc <= 0 {
		for name, command := range godoFile.Commands {
			if command.Description != nil {
				fmt.Printf("%s: %s", name, *command.Description)
			} else {
				fmt.Printf("%s", name)
			}
		}
		return
	}

	command, ok := godoFile.Commands[args[0]]
	if !ok {
		fmt.Printf("%s is not a command in godo.yml\n", args[0])
		return
	}

	if err := engine.Run(command); err != nil {
		fmt.Println(err == nil)
		fmt.Println(err)
	}
}
