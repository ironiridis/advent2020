package main

import (
	"fmt"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

func findTwoSumTo2020(in chan string) (string, error) {
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

func findThreeSumTo2020(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented 😬")
}

func part1() {
	fmt.Println("Day 1, part 1 - expense report values, pairs summing to 2020")
	ans, err := findTwoSumTo2020(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
}

func part2() {
	fmt.Println("Day 1, part 2 - expense report values, triples summing to 2020")
	ans, err := findThreeSumTo2020(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}

func main() {
	part1()
	part2()
}
