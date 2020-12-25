package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestProblemExample(t *testing.T) {
	scando.Strap(findSumTo2020, `1721
		979
		366
		299
		675
		1456`, "514579", t)

}
