package mascot_test

import (
	"testing"

	"example.com/m/v2/01_hello_world/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "go gopher" {
		t.Fatal("Wrong mascot :/")
	}
}
