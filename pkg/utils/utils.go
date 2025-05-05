package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/VincentBrodin/godo/pkg/parser"
)

// Tries to find a file based on name and read it.
// (In short we dont care about extension)
// Name: The name of the file without any path or extension
// Exclude: Any extension that should be ignored, for example .exe or .o
func ReadByName(name string, exclude ...string) ([]byte, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(wd)
	if err != nil {
		return nil, err
	}

outer:
	for _, entrie := range entries {
		if !entrie.Type().IsRegular() || CutExtension(entrie.Name()) != name {
			continue
		}

		extension := GetExtension(entrie.Name())

		for _, ex := range exclude {
			if ex == extension {
				continue outer
			}
		}

		info, err := entrie.Info()
		if err != nil {
			return nil, err
		}

		return os.ReadFile(info.Name())
	}

	return nil, fmt.Errorf("Could not find any file with the name: %s\n", name)
}

func GetExtension(name string) string {
	namec := len(name)
	for _i := range namec {
		i := namec - _i - 1
		if name[i] == '.' {
			return name[i:]
		}
	}
	return name
}

func CutExtension(name string) string {
	namec := len(name)
	for _i := range namec {
		i := namec - _i - 1
		if name[i] == '.' {
			return name[:i]
		}
	}
	return name
}

// Returns recomended defualt running type based on os
func GetDefaultType() string {
	switch runtime.GOOS {
	case "windows":
		return parser.SHELL
	case "linux":
		return parser.RAW
	default:
		return parser.RAW
	}
}

func GetShell() (string, []string) {
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

func GetDir(cmd parser.Command) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return wd, err
	}

	if cmd.Where != nil {
		wd = path.Join(wd, *cmd.Where)
	}
	return wd, nil
}

func GetRunAndType(cmd parser.Command) (parser.Run, *string, error) {
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
