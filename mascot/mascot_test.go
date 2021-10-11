package mascot_test

import (
	"go-trials/mascot"
	"testing"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong Mascot :-(")
	}
}
