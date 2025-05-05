package parser_test

import (
	"testing"

	"github.com/VincentBrodin/godo/pkg/parser"
)

func TestParser(t *testing.T) {
	dummyfile := `
commands:
  test:
    run: 
      - go clean -testcache
      - go test ./...

  os:  
    variants:
      - run: echo Windows
        platform: windows
      - run: echo Linux
        platform: linux
      - run: echo Unknown
        platform: default`

	file, err := parser.Parse([]byte(dummyfile))
	if err != nil {
		t.Logf("Error parsing file: %v\n", err)
		t.FailNow()
	}

	if len(file.Commands) != 2 {
		t.Logf("Unexpected amount of commands: %d expected 2\n", len(file.Commands))
		t.FailNow()
	}

	if len(*file.Commands["test"].Run) != 2 {
		t.Logf("Unexpected amount of run commands in test: %d expected 2\n", len(*file.Commands["test"].Run))
		t.FailNow()
	}

	if len(file.Commands["os"].Variants) != 3 {
		t.Logf("Unexpected amount of variants in os: %d expected 3\n", len(file.Commands["os"].Variants))
		t.FailNow()
	}
}
