package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(findTwoSumTo2020, `1721
		979
		366
		299
		675
		1456`, "514579", t)
}

func TestPart2Example(t *testing.T) {
	scando.Strap(findThreeSumTo2020, `1721
		979
		366
		299
		675
		1456`, "241861950", t)
}
