package main

import (
	"fmt"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

func findSumTo2020(in chan string) (string, error) {
	m := make(map[int]bool)
	for s := range in {
		val, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		if m[2020-val] {
			return strconv.Itoa(val * (2020 - val)), nil
		}
		m[val] = true
	}
	return "", fmt.Errorf("no solution found")
}

func main() {
	fmt.Println("Day 1, part 1 - input expense report values")
	ans, err := findSumTo2020(scando.Stdin())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
}
