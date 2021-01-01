package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

func getlist(in chan string) ([]int, error) {
	adapters := []int{0} // implicit zeroth "joltage"
	for s := range in {
		if s == "" {
			continue
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("cannot convert %q to int: %w", s, err)
		}
		adapters = append(adapters, v)
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	return adapters, nil
}

func part1func(in chan string) (string, error) {
	adapters, err := getlist(in)
	if err != nil {
		return "", fmt.Errorf("cannot get list: %w", err)
	}
	var diff1, diff3 int

	for j := range adapters {
		if j == len(adapters)-1 {
			break
		}
		switch adapters[j+1] - adapters[j] {
		case 1:
			diff1++
		case 3:
			diff3++
		default:
			return "", fmt.Errorf("between adapters[%d] and adapters[%d] there's an unhandled difference of %d", j, j+1, adapters[j+1]-adapters[j])
		}
	}
	return strconv.Itoa(diff1 * diff3), nil
}

func part2func(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 10, part 1 - product of 1-diff and 3-diff jolts in list")
	ans, err := part1func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 10, part 2 - summary")
	ans, err = part2func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
