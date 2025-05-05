package engine

import (
	"os"
	"os/exec"

	"github.com/VincentBrodin/godo/pkg/utils"
	"github.com/google/shlex"
)

func runShell(resCmd ResolvedCommand) error {
	for _, run := range resCmd.Run {
		// EMPTY
		if len(run) == 0 {
			continue
		}
		// Check if is banglines
		canFail := false
		if run[0] == '$' {
			canFail = true
			run = run[1:]
		}

		split, err := shlex.Split(run)
		if err != nil {
			return err
		}

		shell, args := utils.GetShell()

		args = append(args, split...)
		cmd := exec.Command(shell, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = resCmd.Where
		if err := cmd.Run(); err != nil && !canFail {
			return err
		}

	}
	return nil

}

func runPath(resCmd ResolvedCommand) error {
	for _, run := range resCmd.Run {
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
		cmd.Dir = resCmd.Where
		if err := cmd.Run(); err != nil {
			return err
		}

	}

	return nil
}

func runRaw(resCmd ResolvedCommand) error {
	for _, run := range resCmd.Run {
		split, err := shlex.Split(run)
		if err != nil {
			return err
		}

		cmd := exec.Command(split[0], split[1:]...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = resCmd.Where

		if err := cmd.Run(); err != nil {
			return err
		}

	}
	return nil
}
