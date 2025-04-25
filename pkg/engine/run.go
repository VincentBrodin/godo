package engine

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/VincentBrodin/godo/pkg/parser"
	"github.com/google/shlex"
)

// Run the command
func Run(cmd parser.Command) error {
	t := ""
	if cmd.Type == nil {
		t = "path"
	} else {
		t = *cmd.Type
	}
	for _, run := range cmd.Run {
		switch t {
		case "raw":
			if err := rawRun(cmd, run); err != nil {
				return err
			}
		case "path":
			if err := pathRun(cmd, run); err != nil {
				return err
			}
		case "shell":
			if err := shellRun(cmd, run); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Unkown run type: %s. Only allows: 'raw','path' or 'shell' (defualt 'path')\n")
		}
	}
	return nil
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

func pathRun(command parser.Command, run string) error {
	split, err := shlex.Split(run)
	if err != nil {
		return err
	}

	path, err := exec.LookPath(split[0])
	if err != nil {
		return err
	}

	cmd := exec.Command(path, split[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if wd, err := getDir(command); err == nil {
		cmd.Dir = wd
	}
	if command.Wait {
		return cmd.Start()
	} else {
		return cmd.Run()
	}
}

func rawRun(command parser.Command, run string) error {
	split, err := shlex.Split(run)
	if err != nil {
		return err
	}
	cmd := exec.Command(split[0], split[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if wd, err := getDir(command); err == nil {
		cmd.Dir = wd
	}
	if command.Wait {
		return cmd.Start()
	} else {
		return cmd.Run()
	}

}

func shellRun(command parser.Command, run string) error {
	split, err := shlex.Split(run)
	if err != nil {
		return err
	}

	var shell string
	var args []string

	if runtime.GOOS == "windows" {
		shell = "cmd"
		args = []string{"/C"}
	} else {
		shell = "sh"
		args = []string{"-c"}
	}

	args = append(args, split...)
	cmd := exec.Command(shell, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if wd, err := getDir(command); err == nil {
		cmd.Dir = wd
	}
	if command.Wait {
		return cmd.Start()
	} else {
		return cmd.Run()
	}

}
