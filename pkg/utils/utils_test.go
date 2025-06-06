package utils_test

import (
	"github.com/VincentBrodin/godo/pkg/utils"
	"testing"
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

func TestCanFail(t *testing.T) {
	expected := "test"
	inputA := "$test"
	inputB := "test"

	if str, canFail := utils.CanFail(inputA); str != expected || !canFail {
		t.Logf("Input A failed")
		t.FailNow()
	}

	if str, canFail := utils.CanFail(inputB); str != expected || canFail {
		t.Logf("Input B failed")
		t.FailNow()
	}
}
