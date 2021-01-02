package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(part1func,
		`F10
		N3
		F7
		R90
		F11`, "25", t)
}

func TestPart2Example(t *testing.T) {
	scando.Strap(part2func,
		`F10
		N3
		F7
		R90
		F11`, "286", t)
}
