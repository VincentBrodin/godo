package engine

import (
	"testing"

	"github.com/VincentBrodin/godo/pkg/parser"
)

func TestNoVariants(t *testing.T) {
	expected := "expected"
	unexpected := "unexpected"
	a := new(parser.Run)
	a.Add(expected)
	var b parser.Run
	b.Add(unexpected)
	cmd := parser.Command{
		Run: a,
		Variants: []parser.Variant{
			{
				Run: b,
			},
		},
	}

	resCmd, err := resolve(cmd)

	if err != nil {
		t.Logf("Got and unexpected error: %v\n", err)
		t.FailNow()
	}

	if resCmd.Run[0] != expected {
		t.Logf("Got and unexpected value: %s\n", resCmd.Run[0])
		t.FailNow()
	}
}

func TestVariants(t *testing.T) {
	expected := "expected"
	var a parser.Run
	a.Add(expected)
	cmd := parser.Command{
		Variants: []parser.Variant{
			{
				Run: a,
				Platform: "defualt",
			},
		},
	}

	resCmd, err := resolve(cmd)

	if err != nil {
		t.Logf("Got and unexpected error: %v\n", err)
		t.FailNow()
	}

	if resCmd.Run[0] != expected {
		t.Logf("Got and unexpected value: %s\n", resCmd.Run[0])
		t.FailNow()
	}
}
