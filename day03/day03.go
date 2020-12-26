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
func (m treemap) treeAt(x, y int) (bool, error) {
	if x < 0 {
		// probably could support this, but assume we won't
		return false, fmt.Errorf("x coord %d is negative", x)
	}
	if y < 0 {
		return false, fmt.Errorf("y coord %d is negative", y)
	}
	if y >= len(m) {
		return false, fmt.Errorf("y coord %d exceeds rows in map (0-%d)", y, len(m)-1)
	}
	// this assumes all rows are the same width, which must be true for horizontal repeating to happen
	return '#' == m[y][x%len(m[0])], nil
}
func (m treemap) checkSlope(xincr, yincr int) (int, error) {
	if yincr <= 0 {
		return 0, fmt.Errorf("y slope of %d invalid, must be at least 1", yincr)
	}
	var encounters int
	var x, y int
	for {
		isTree, err := m.treeAt(x, y)
		if err != nil {
			return 0, err
		}
		if isTree {
			encounters++
		}
		x += xincr
		y += yincr
		if y >= m.rows() {
			break
		}
	}
	return encounters, nil
}

func countTreeEncountersPart1(in chan string) (string, error) {
	m, err := makeTreeMap(in)
	if err != nil {
		return "", err
	}
	encounters, err := m.checkSlope(3, 1)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(encounters), nil
}

func countTreeEncountersPart2(in chan string) (string, error) {
	var encounters int
	m, err := makeTreeMap(in)
	if err != nil {
		return "", err
	}
	for _, c := range [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		r, err := m.checkSlope(c[0], c[1])
		if err != nil {
			return "", err
		}
		if encounters == 0 {
			encounters = r
		} else {
			encounters *= r
		}
	}
	return strconv.Itoa(encounters), nil
}

func main() {
	fmt.Println("Day 3, part 1 - tree encounter count")
	ans, err := countTreeEncountersPart1(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 3, part 2 - tree encounter count, several scenarios")
	ans, err = countTreeEncountersPart2(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
