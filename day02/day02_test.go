package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(countValidPasswords, `1-3 a: abcde
	1-3 b: cdefg
	2-9 c: ccccccccc`, "2", t)
}

func TestPart2Example(t *testing.T) {
	t.Fail()

}
