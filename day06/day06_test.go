package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(part1func, `abc

	a
	b
	c
	
	ab
	ac
	
	a
	a
	a
	a
	
	b`, "11", t)
}

func TestPart2Example(t *testing.T) {
	scando.Strap(part2func, `abc

	a
	b
	c
	
	ab
	ac
	
	a
	a
	a
	a
	
	b`, "6", t)
}
