package main

import (
	"fmt"

	"github.com/ironiridis/advent2020/scando"
)

func countValidPasswords(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 2, part 1 - expense report values, pairs summing to 2020")
	ans, err := countValidPasswords(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
}
