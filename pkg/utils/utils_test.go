package utils_test

import (
	"testing"
	"github.com/VincentBrodin/godo/pkg/utils"
)

func TestCutExtension(t *testing.T) {
	input := "test.go"
	expected := "test"
	output := utils.CutExtension(input)
	t.Logf("Sent %s expects %s got %s\n", input, expected, output)
	if output != expected {
		t.FailNow()
	}
}

func TestGetExtension(t *testing.T) {
	input := "test.go"
	expected := ".go"
	output := utils.GetExtension(input)
	t.Logf("Sent %s expects %s got %s\n", input, expected, output)
	if output != expected {
		t.FailNow()
	}
}
