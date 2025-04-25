package engine

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/VincentBrodin/godo/pkg/parser"
	"github.com/google/shlex"
)

func Run(command parser.Command) error {
	var totErrs error
	for _, run := range command.Run {
		tempErr := rawRun(command, run)
		if tempErr == nil {
			continue
		}
		errs := errors.Join(tempErr)

		tempErr = pathRun(command, run)
		if tempErr == nil {
			continue
		}
		errs = errors.Join(tempErr)

		tempErr = shellRun(command, run)
		if tempErr == nil {
			continue
		}
		errs = errors.Join(tempErr)
		totErrs = errors.Join(errs)
	}

	return totErrs
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
	if wd, err := getDir(command); err != nil {
		cmd.Dir = wd
	}
	return cmd.Run()
}

func rawRun(command parser.Command, run string) error {
	split, err := shlex.Split(run)
	if err != nil {
		return err
	}
	cmd := exec.Command(split[0], split[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if wd, err := getDir(command); err != nil {
		cmd.Dir = wd
	}
	return cmd.Run()
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
	if wd, err := getDir(command); err != nil {
		cmd.Dir = wd
	}
	return cmd.Run()
}
