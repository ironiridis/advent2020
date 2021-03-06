package main

import (
	"fmt"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

type answers map[rune]int

func readGroup(in chan string) (a answers, folks int, done bool, err error) {
	a = make(answers, 0)
	for {
		s, ok := <-in
		if !ok {
			done = true
		}
		if s == "" {
			return
		}
		folks++
		for _, sym := range s {
			a[sym]++
		}
	}
}

func part1func(in chan string) (string, error) {
	var total int
	for {
		a, _, done, err := readGroup(in)
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
	var total int
	for {
		a, folks, done, err := readGroup(in)
		if err != nil {
			return "", err
		}
		for sym := range a {
			if a[sym] == folks {
				total++
			}
		}
		if done {
			break
		}
	}
	return strconv.Itoa(total), nil
}

func main() {
	fmt.Println("Day 6, part 1 - sum of group affirmative answers")
	ans, err := part1func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 6, part 2 - sum of unanimous affirmative answers")
	ans, err = part2func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
