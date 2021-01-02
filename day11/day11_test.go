package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(part1func,
		`L.LL.LL.LL
		 LLLLLLL.LL
		 L.L.L..L..
		 LLLL.LL.LL
		 L.LL.LL.LL
		 L.LLLLL.LL
		 ..L.L.....
		 LLLLLLLLLL
		 L.LLLLLL.L
		 L.LLLLL.LL`, "37", t)
}

func TestPart2Example(t *testing.T) {
	scando.Strap(part2func,
		`L.LL.LL.LL
		 LLLLLLL.LL
		 L.L.L..L..
		 LLLL.LL.LL
		 L.LL.LL.LL
		 L.LLLLL.LL
		 ..L.L.....
		 LLLLLLLLLL
		 L.LLLLLL.L
		 L.LLLLL.LL`, "26", t)
}
