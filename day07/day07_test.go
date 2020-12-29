package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(part1func, "example data", "expected output", t)
}

func TestPart2Example(t *testing.T) {
	t.Fail()
}
