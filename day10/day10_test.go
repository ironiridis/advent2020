package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(part1func, `16
	10
	15
	5
	1
	11
	7
	19
	6
	12
	4`, "35", t)

	scando.Strap(part1func, `28
	33
	18
	42
	31
	14
	46
	20
	48
	47
	24
	23
	49
	45
	19
	38
	39
	11
	1
	32
	25
	35
	8
	17
	7
	9
	4
	2
	34
	10
	3`, "220", t)

}

func TestPart2Example(t *testing.T) {
	scando.Strap(part2func, `16
	10
	15
	5
	1
	11
	7
	19
	6
	12
	4`, "8", t)

	scando.Strap(part2func, `28
	33
	18
	42
	31
	14
	46
	20
	48
	47
	24
	23
	49
	45
	19
	38
	39
	11
	1
	32
	25
	35
	8
	17
	7
	9
	4
	2
	34
	10
	3`, "19208", t)

}
