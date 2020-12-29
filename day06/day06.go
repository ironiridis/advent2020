package main

import (
	"fmt"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

type answers map[rune]int

func readGroup(in chan string) (a answers, done bool, err error) {
	a = make(answers, 0)
	for {
		s, ok := <-in
		if !ok {
			done = true
		}
		if s == "" {
			return
		}
		for _, sym := range s {
			a[sym]++
		}
	}
}

func part1func(in chan string) (string, error) {
	var total int
	for {
		a, done, err := readGroup(in)
		if err != nil {
			return "", err
		}
		total += len(a)
		if done {
			break
		}
	}
	return strconv.Itoa(total), nil
}

func part2func(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 6, part 1 - sum of group affirmative answers")
	ans, err := part1func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 6, part 2 - summary")
	ans, err = part2func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
