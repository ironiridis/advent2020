package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(part1func, `939
	7,13,x,x,59,x,31,19`, "295", t)
}

func TestPart2Example(t *testing.T) {
	t.Fail()
}
