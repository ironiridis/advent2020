package main

import (
	"fmt"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

func seatId(s string) (int, error) {
	if len(s) != 10 {
		return 0, fmt.Errorf("input %q is not 10 characters", s)
	}
	pval := 512
	sid := 0
	for _, c := range s {
		if c == 'B' || c == 'R' {
			sid += pval
		}
		pval = pval >> 1
	}
	return sid, nil
}

func getMaximumSeatID(in chan string) (string, error) {
	var max int
	for s := range in {
		if s == "" { // ignore insignificant blank lines
			continue
		}
		sid, err := seatId(s)
		if err != nil {
			return "", err
		}
		if sid > max {
			max = sid
		}
	}
	return strconv.Itoa(max), nil
}

func getMissingSeatID(in chan string) (string, error) {
	var seats [1024]bool
	for s := range in {
		if s == "" { // ignore insignificant blank lines
			continue
		}
		sid, err := seatId(s)
		if err != nil {
			return "", err
		}
		seats[sid] = true
	}

	for p := range seats {
		if p == 0 {
			continue
		}
		if p == len(seats)-1 {
			break
		}
		if seats[p-1] && !seats[p] && seats[p+1] {
			return strconv.Itoa(p), nil
		}
	}
	return "", fmt.Errorf("did not find an unset seat id")
}

func main() {
	fmt.Println("Day 5, part 1 - highest seat id")
	ans, err := getMaximumSeatID(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 5, part 2 - missing seat id")
	ans, err = getMissingSeatID(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
