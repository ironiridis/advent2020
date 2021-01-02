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

func part2walk(tree map[int][]int, tgt int) int64 {
	for {
		// 0 is the goal joltage; if we reach that, return 1
		if tgt == 0 {
			return 1
		}
		switch len(tree[tgt]) {
		case 0:
			// shouldn't be possible, but here for completeness
			return 0
		case 1:
			// rather than recursing or creating new goroutines, just reuse this goroutine
			tgt = tree[tgt][0]
			continue
		default:
			// calculate depth first inline
			res := part2walk(tree, tree[tgt][0])
			gather := make(chan int64)
			for _, next := range tree[tgt][1:] {
				go func(next int, out chan int64) {
					out <- part2walk(tree, next)
				}(next, gather)
			}
			for n := len(tree[tgt]) - 1; n > 0; n-- {
				res += <-gather
			}
			close(gather)
			return res
		}
	}
}

func part2func(in chan string) (string, error) {
	list, err := getlist(in)
	if err != nil {
		return "", fmt.Errorf("cannot get list: %w", err)
	}

	tree := make(map[int][]int, len(list))
	for j := len(list) - 1; j >= 0; j-- {
		tree[list[j]] = make([]int, 0, 3)
		for v := range tree {
			if v == j {
				continue
			}
			if n := v - list[j]; n <= 3 && n >= 1 {
				tree[v] = append(tree[v], list[j])
			}
		}
	}

	combinations := part2walk(tree, list[len(list)-1])
	return strconv.FormatInt(combinations, 10), nil
}

func main() {
	fmt.Println("Day 10, part 1 - product of 1-diff and 3-diff jolts in list")
	ans, err := part1func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 10, part 2 - distinct combinations of adapters")
	ans, err = part2func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
