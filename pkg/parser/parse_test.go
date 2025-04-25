package parser_test

import (
	"testing"

	"github.com/VincentBrodin/godo/pkg/parser"
)

// Test to see if Run type correctly unmarshals scalar types
func TestRunScalar(t *testing.T) {
	file := `
commands:
  test:
    run: go run .
`
	godo, err := parser.Parse([]byte(file))
	if err != nil {
		t.Logf("Failed to parse file: %s", err)
		t.FailNow()
	}
	cmd, ok := godo.Commands["test"]
	if !ok {
		t.Log("Failed to find command")
		t.FailNow()
	}

	if len(cmd.Run) != 1 {
		t.Log("Unexpected amount of runs")
		t.FailNow()
	}
}

// Test to see if Run type correctly unmarshals sequence types
func TestRunSequence(t *testing.T) {
	file := `
commands:
  test:
    run:
      - go test
      - go run .
`
	godo, err := parser.Parse([]byte(file))
	if err != nil {
		t.Logf("Failed to parse file: %s", err)
		t.FailNow()
	}

	cmd, ok := godo.Commands["test"]
	if !ok {
		t.Log("Failed to find command")
		t.FailNow()
	}

	if len(cmd.Run) != 2 {
		t.Log("Unexpected amount of runs")
		t.FailNow()
	}
}
