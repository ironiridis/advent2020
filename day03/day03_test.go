package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(countTreeEncounters,
		`..##.......
		 #...#...#..
		 .#....#..#.
		 ..#.#...#.#
		 .#...##..#.
		 ..#.##.....
		 .#.#.#....#
		 .#........#
		 #.##...#...
		 #...##....#
		 .#..#...#.#`, "7", t)
}

func TestPart2Example(t *testing.T) {
	t.Fail()
}
