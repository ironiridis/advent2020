package main

import (
	"fmt"

	"github.com/ironiridis/advent2020/scando"
)

func part1func(in chan string, n int) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func part2func(in chan string, n int) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 9, part 1 - summary")
	ans, err := part1func(scando.Input(), 25)
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 9, part 2 - summary")
	ans, err = part2func(scando.Input(), 25)
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
