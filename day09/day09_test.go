package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func part1funcT(in chan string) (string, error) {
	return part1func(in, 5)
}

func part2funcT(in chan string) (string, error) {
	return part2func(in, 5)
}

func TestPart1Example(t *testing.T) {
	scando.Strap(part1funcT, `35
	20
	15
	25
	47
	40
	62
	55
	65
	95
	102
	117
	150
	182
	127
	219
	299
	277
	309
	576`, "127", t)
}

func TestPart2Example(t *testing.T) {
	scando.Strap(part2funcT, `35
	20
	15
	25
	47
	40
	62
	55
	65
	95
	102
	117
	150
	182
	127
	219
	299
	277
	309
	576`, "62", t)
}
