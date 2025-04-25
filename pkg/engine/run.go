package engine

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/VincentBrodin/godo/pkg/parser"
	"github.com/google/shlex"
)

// Run the command
func Run(cmd parser.Command) error {
	var t string
	if cmd.Type == nil {
		t = "shell"
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
			return fmt.Errorf("Unkown run type: %s. Only allows: 'raw','path' or 'shell' (defualt 'shell')\n", t)
		}
	}
	return nil
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

	shell, args := getShell()

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
