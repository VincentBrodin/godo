package engine

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/VincentBrodin/godo/pkg/parser"
)

// Returns recomended defualt running type based on os
func getDefaultType() string {
	switch runtime.GOOS {
	case "windows":
		return parser.SHELL
	case "linux":
		return parser.RAW
	default:
		return parser.RAW
	}
}

func getShell() (string, []string) {
	var shell string
	var args []string
	if runtime.GOOS == "windows" {
		if pwsh, err := exec.LookPath("pwsh"); err == nil {
			shell = pwsh
			args = []string{"-NoProfile", "-Command"}
		} else {
			shell = "cmd"
			args = []string{"/C"}
		}
	} else {
		if bash, err := exec.LookPath("bash"); err == nil {
			shell = bash
		} else {
			shell = "sh"
		}
		args = []string{"-c"}
	}
	return shell, args
}

func getDir(cmd parser.Command) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return wd, err
	}

	if cmd.Where != nil {
		wd = path.Join(wd, *cmd.Where)
	}
	return wd, nil
}

func getRunAndType(cmd parser.Command) (parser.Run, *string, error) {
	if cmd.Run == nil {
		if len(cmd.Variants) == 0 {
			return nil, nil, fmt.Errorf("Command has nothing to run\n")
		}

		// Check for platform
		for _, variant := range cmd.Variants {
			if variant.Platform == runtime.GOOS {
				return variant.Run, variant.Type, nil
			}
		}

		// Check for fallback
		for _, variant := range cmd.Variants {
			if variant.Platform == "defualt" {
				return variant.Run, variant.Type, nil
			}
		}

		return nil, nil, fmt.Errorf("Could not find a variant that matches system os (%s)\n", runtime.GOOS)
	} else {
		return *cmd.Run, cmd.Type, nil
	}

}
