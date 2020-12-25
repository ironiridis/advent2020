package main

import (
	"flag"
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
	fmt.Println("Day 1, part 1 - input expense report values")
	ans, err := findTwoSumTo2020(scando.Stdin())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
}

func part2() {
	fmt.Println("Day 1, part 2 - input expense report values again")
	ans, err := findThreeSumTo2020(scando.Stdin())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Which Part of the challenge to run")
	flag.Parse()
	switch part {
	case 1:
		part1()
	case 2:
		part2()
	default:
		panic(fmt.Errorf("Cannot run part %d", part))
	}
}
