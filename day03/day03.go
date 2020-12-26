package main

import (
	"fmt"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

type treemap [][]byte

func makeTreeMap(in chan string) (treemap, error) {
	var m treemap
	for s := range in {
		m = append(m, []byte(s))
		if len(s) != len(m[0]) {
			return nil, fmt.Errorf("length of %q does not equal length of first row (%d)", s, len(m[0]))
		}
	}
	return m, nil
}
func (m treemap) rows() int { return len(m) }
func (m treemap) get(x, y int) (byte, error) {
	if x < 0 {
		// probably could support this, but assume we won't
		return 0, fmt.Errorf("x coord %d is negative", x)
	}
	if y < 0 {
		return 0, fmt.Errorf("y coord %d is negative", y)
	}
	if y >= len(m) {
		return 0, fmt.Errorf("y coord %d exceeds rows in map (0-%d)", y, len(m)-1)
	}
	// this assumes all rows are the same width, which must be true for horizontal repeating to happen
	return m[y][x%len(m[0])], nil
}

func countTreeEncounters(in chan string) (string, error) {
	xincr, yincr := 3, 1
	encounters := 0
	m, err := makeTreeMap(in)
	if err != nil {
		return "", err
	}
	var x, y int
	for {
		cell, err := m.get(x, y)
		if err != nil {
			return "", err
		}
		if cell == '#' {
			encounters++
		}
		x += xincr
		y += yincr
		if y >= m.rows() {
			break
		}
	}
	return strconv.Itoa(encounters), nil
}

func main() {
	fmt.Println("Day 3, part 1 - tree encounter count")
	ans, err := countTreeEncounters(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	/*
		fmt.Println("Day 3, part 2 - password validity, mutually exclusive symbol positions")
		ans, err = alternateCountValidPasswords(scando.Input())
		if err != nil {
			fmt.Printf("Cannot determine answer: %v\n", err)
			return
		}
		fmt.Printf("Part 2 Answer: %q\n", ans)
	*/
}
