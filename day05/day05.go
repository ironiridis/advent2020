package main

import (
	"fmt"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

func getMaximumSeatID(in chan string) (string, error) {
	var max int
	for s := range in {
		if s == "" { // ignore insignificant blank lines
			continue
		}
		if len(s) != 10 {
			return "", fmt.Errorf("input %q is not 10 bytes", s)
		}
		pval := 512
		sid := 0
		for _, c := range s {
			if c == 'B' || c == 'R' {
				sid += pval
			}
			pval = pval >> 1
		}
		if sid > max {
			max = sid
		}
	}
	return strconv.Itoa(max), nil
}

func main() {
	fmt.Println("Day 5, part 1 - highest seat id")
	ans, err := getMaximumSeatID(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	/*
		fmt.Println("Day 5, part 2 ...")
		ans, err = getMaximumSeatID(scando.Input())
		if err != nil {
			fmt.Printf("Cannot determine answer: %v\n", err)
			return
		}
		fmt.Printf("Part 2 Answer: %q\n", ans)
	*/
}
