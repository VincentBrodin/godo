package engine

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/VincentBrodin/godo/pkg/parser"
	"github.com/google/shlex"
)

func Run(command parser.Command) {
	for _, run := range command.Run {
		if err := rawRun(run); err == nil {
			break
		}

		if err := pathRun(run); err == nil {
			break
		}

		if err := shellRun(run); err == nil {
			break
		}
	}
}

func pathRun(run string) error {
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
	return cmd.Run()
}

func rawRun(run string) error {
	split, err := shlex.Split(run)
	if err != nil {
		return err
	}
	cmd := exec.Command(split[0], split[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func shellRun(run string) error {
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
	return cmd.Run()
}
