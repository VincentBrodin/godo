package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/VincentBrodin/godo/pkg/engine"
	"github.com/VincentBrodin/godo/pkg/parser"
	"github.com/VincentBrodin/godo/pkg/utils"
	"github.com/VincentBrodin/whale/codes"
	"github.com/VincentBrodin/whale/list"

	"slices"

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
				if args[0] == "init" {
					wd, err := os.Getwd()
					if err != nil {
						os.Exit(5)
					}

					name := path.Join(wd, "godo.yml")
					if err := os.WriteFile(name, []byte(Init), 0664); err != nil {
						log.Printf("Problem creating godo file: %v\n", err)
						os.Exit(6)
					}
					log.Printf("Created godo file @ %s\n", name)
					os.Exit(0)
				} else {
					log.Printf("Problem reading godo file: %v\n", err)
					os.Exit(4)
				}
			}

			if len(args) == 0 {
				if err := listCommands(godoFile); err != nil {
					os.Exit(2)
				}
				return
			}

			command, ok := godoFile.Commands[args[0]]
			if !ok {
				log.Printf("%s is not a command in godo file\n", args[0])
				os.Exit(3)
			}

			if err := engine.Run(command); err != nil {
				// fmt.Println(err)
				os.Exit(2)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Printf("Got an error: %v\n", err)
		// os.Exit(1)
	}
}

func loadFile() (*parser.GodoFile, error) {
	file, err := utils.ReadByName("godo", ".exe", ".exe~", ".dll", ".so", ".dylib", ".test", ".out")
	if err != nil {
		return nil, err
	}

	godoFile, err := parser.Parse(file)
	if err != nil {
		return nil, err
	}

	return godoFile, nil
}

func listCommands(godoFile *parser.GodoFile) error {
	keys := make([]string, 0, len(godoFile.Commands))
	for k := range godoFile.Commands {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	keys = append(keys, "close")

	l := list.New(list.DefualtConfig())
	l.Config.RenderItem = func(item string, selected bool, config list.Config) string {
		description := ""
		if godoFile.Commands[item].Description != nil {
			description = *godoFile.Commands[item].Description
		}
		if selected {
			return fmt.Sprintf("  > %s: %s%s", item, codes.Muted, description)
		}
		return fmt.Sprintf("%s    %s: %s", codes.Muted, item, description)
	}
	l.Config.Lable = "Select command"

	index, err := l.Prompt(keys)

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return err
	}

	if index == len(keys)-1 {
		return nil
	} else {
		if err := engine.Run(godoFile.Commands[keys[index]]); err != nil {
			return err
		}
	}
	return nil
}
