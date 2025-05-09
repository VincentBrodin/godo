package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/VincentBrodin/godo/pkg/engine"
	"github.com/VincentBrodin/godo/pkg/parser"
	"github.com/VincentBrodin/godo/pkg/utils"
	"github.com/VincentBrodin/suddig/matcher"

	"slices"

	"github.com/manifoldco/promptui"
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
	keys := make([]string, 0, len(godoFile.Commands))
	for k := range godoFile.Commands {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	keys = append(keys, "close")

	prompt := promptui.Select{
		Label: "Select command",
		Searcher: func(input string, index int) bool {
			c := matcher.DefualtConfig()
			c.MinScore = 0.5
			c.StringFunc = func(s string) string { return strings.ToLower(s) }
			m := matcher.New(c)
			return m.Match(input, keys[index]) || strings.HasPrefix(keys[index], input)
		},
		Size:  len(keys),
		Items: keys,
	}
	index, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if index == len(keys)-1 {
		return
	} else {
		if err := engine.Run(godoFile.Commands[result]); err != nil {
			fmt.Println(err)
		}

	}
}
