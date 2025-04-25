package engine

import (
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/VincentBrodin/godo/pkg/parser"
)

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

func getDir(command parser.Command) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return wd, err
	}

	if command.Where != nil {
		wd = path.Join(wd, *command.Where)
	}
	return wd, nil
}
