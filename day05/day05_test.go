package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(getMaximumSeatID, `FFFFFFFLLL`, "0", t)
	scando.Strap(getMaximumSeatID, `BBBBBBBRRR`, "1023", t)

	scando.Strap(getMaximumSeatID, `BFFFBBFRRR`, "567", t)
	scando.Strap(getMaximumSeatID, `FFFBBBFRRR`, "119", t)
	scando.Strap(getMaximumSeatID, `BBFFBBFRLL`, "820", t)

	scando.Strap(getMaximumSeatID,
		`FFFFFFFLLL
		FFFFFFFLLR
		FFFFFFFLRL
		`, "2", t)
}

func TestPart2Example(t *testing.T) {
	scando.Strap(getMaximumSeatID,
		`FFFFFFFLRL
		FFFFFFFRLL
		`, "4", t)
}
