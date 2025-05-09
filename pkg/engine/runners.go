package engine

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/VincentBrodin/godo/pkg/utils"
	"github.com/google/shlex"
)

func runShell(resCmd ResolvedCommand) error {
	for _, run := range resCmd.Run {
		// EMPTY
		run = strings.TrimSpace(run)
		if len(run) == 0 {
			continue
		}
		// Check if is banglines
		run, canFail := utils.CanFail(run)

		if canFail {
			log.Printf("Running %s (can fail)\n", run)
		} else {
			log.Printf("Running %s\n", run)
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

		if err := runCmd(cmd); err != nil {
			if canFail {
				log.Printf("Got an error running '%s', but will keep running: %v\n", run, err)
			} else {
				return err
			}
		}

	}
	return nil

}

func runPath(resCmd ResolvedCommand) error {

	for _, run := range resCmd.Run {
		// EMPTY
		run = strings.TrimSpace(run)
		if len(run) == 0 {
			continue
		}
		// Check if is banglines
		run, canFail := utils.CanFail(run)

		if canFail {
			log.Printf("Running %s (can fail)\n", run)
		} else {
			log.Printf("Running %s\n", run)
		}

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

		if err := runCmd(cmd); err != nil {
			if canFail {
				log.Printf("Got an error running '%s', but will keep running: %v\n", run, err)
			} else {
				return err
			}
		}

	}

	return nil
}

func runRaw(resCmd ResolvedCommand) error {
	for _, run := range resCmd.Run {
		// EMPTY
		run = strings.TrimSpace(run)
		if len(run) == 0 {
			continue
		}
		run, canFail := utils.CanFail(run)

		if canFail {
			log.Printf("Running %s (can fail)\n", run)
		} else {
			log.Printf("Running %s\n", run)
		}

		split, err := shlex.Split(run)
		if err != nil {
			return err
		}

		cmd := exec.Command(split[0], split[1:]...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = resCmd.Where

		if err := runCmd(cmd); err != nil {
			if canFail {
				log.Printf("Got an error running '%s', but will keep running: %v\n", run, err)
			} else {
				return err
			}
		}

	}
	return nil
}

func runCmd(cmd *exec.Cmd) error {
	if err := cmd.Run(); err != nil {
		return err
	} else if exit := cmd.ProcessState.ExitCode(); exit != 0 {
		return fmt.Errorf("Exited with status code %d\n", exit)
	}

	return nil
}
