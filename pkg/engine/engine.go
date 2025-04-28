package engine

import (
	"fmt"

	"github.com/VincentBrodin/godo/pkg/parser"
)

// Run the command
func Run(cmd parser.Command) error {
	resCmd, err := resolve(cmd)
	if err != nil {
		return err
	}

	switch resCmd.Type {
	case "raw":
		if err := runRaw(resCmd); err != nil {
			return err
		}
	case "path":
		if err := runPath(resCmd); err != nil {
			return err
		}
	case "shell":
		if err := runShell(resCmd); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unkown run type: %s. Only allows: 'raw','path' or 'shell' (defualt 'shell')\n", resCmd.Type)
	}

	return nil
}
