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

type treeType map[int][]int

func part2walk(tree treeType, tgt int) int64 {
	if tgt == -1 {
		for k := range tree {
			if k > tgt {
				tgt = k
			}
		}
	}
	for {
		switch len(tree[tgt]) {
		case 0:
			// we reached the root of this tree
			return 1
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

func treeFromList(list []int) treeType {
	tree := make(treeType, len(list))
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
	return tree
}

func part2func(in chan string) (string, error) {
	list, err := getlist(in)
	if err != nil {
		return "", fmt.Errorf("cannot get list: %w", err)
	}
	var trees []treeType
	var lastCut int
	for idx := range list {
		if idx == 0 {
			continue
		}
		if list[idx]-list[idx-1] == 3 {
			//fmt.Printf("%d: cut %d-%d: %d\n", list, lastCut, idx, list[lastCut:idx+1])
			trees = append(trees, treeFromList(list[lastCut:idx+1]))
			lastCut = idx
		}
	}
	//fmt.Printf("%#v\n", trees)
	combinations := int64(1)
	for idx := range trees {
		r := part2walk(trees[idx], -1)
		fmt.Printf("trees[%d]=%#v; r=%d\n", idx, trees[idx], r)
		combinations *= r
	}
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
