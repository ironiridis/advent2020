package main

import (
	"fmt"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

func isSumOfTwoOf(v int64, r []int64) bool {
	// this is a very bad implementation; at a minimum we should only scan k
	// starting at j+1, but the sets for these are always going to be very
	// small (sets of 25 -> 600 comparisons) so i'll keep it simple/obvious
	// if the sets were large, this would be worth revisiting
	for j := range r {
		for k := range r {
			if j == k {
				continue
			}
			if r[j]+r[k] == v {
				return true
			}
		}
	}
	return false
}

func findBadNumber(in chan string, n int) (int64, []int64, error) {
	var list []int64
	for s := range in {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, nil, fmt.Errorf("cannot parse %q as int64: %w", s, err)
		}
		list = append(list, v)
		// preamble phase
		if len(list) <= n {
			continue
		}
		// go's slice index calculation makes me double-take every single time
		if !isSumOfTwoOf(v, list[(len(list)-n)-1:len(list)-1]) {
			return v, list, nil
		}
	}
	return 0, list, fmt.Errorf("did not find candidate value")
}

func part1func(in chan string, n int) (string, error) {
	v, _, err := findBadNumber(in, n)
	return strconv.FormatInt(v, 10), err
}

func part2func(in chan string, n int) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 9, part 1 - member number not a sum of prior distinct two n")
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
