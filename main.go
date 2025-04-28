package main

import (
	"fmt"
	"os"

	"github.com/VincentBrodin/godo/pkg/engine"
	"github.com/VincentBrodin/godo/pkg/parser"
	"github.com/VincentBrodin/godo/pkg/utils"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "godo [task]",
		Short: "Godo is a task runner",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			godoFile, err := loadFile()
			if err != nil {
				return
			}

			if len(args) == 0 {
				listCommands(godoFile)
				return
			}

			command, ok := godoFile.Commands[args[0]]
			if !ok {
				fmt.Printf("%s is not a command in godo file\n", args[0])
				return
			}

			if err := engine.Run(command); err != nil {
				fmt.Println(err)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadFile() (*parser.GodoFile, error) {
	file, err := utils.ReadByName("godo", ".exe", ".exe~", ".dll", ".so", ".dylib", ".test", ".out")
	if err != nil {
		fmt.Println("No godo file found!")
		return nil, err
	}

	godoFile, err := parser.Parse(file)
	if err != nil {
		fmt.Printf("Could not parse godo file: %s", err)
		return nil, err
	}

	return godoFile, nil
}

func listCommands(godoFile *parser.GodoFile) {
	for name, command := range godoFile.Commands {
		if command.Description != nil {
			fmt.Printf("%s:\n  -  %s\n", name, *command.Description)
		} else {
			fmt.Printf("%s\n", name)
		}
	}
}
