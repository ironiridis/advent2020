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
	scando.Strap(part2func, `0
	7,13,x,x,59,x,31,19`, "1068781", t)
	scando.Strap(part2func, `0
	17,x,13,19`, "3417", t)
	scando.Strap(part2func, `0
	67,7,59,61`, "754018", t)
	scando.Strap(part2func, `0
	67,x,7,59,61`, "779210", t)
	scando.Strap(part2func, `0
	67,7,x,59,61`, "1261476", t)
	scando.Strap(part2func, `0
	1789,37,47,1889`, "1202161486", t)

}
