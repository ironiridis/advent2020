package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ironiridis/advent2020/scando"
)

func part1func(in chan string) (string, error) {
	var err error
	var startTime int
	var busIDs []int
	if s, ok := <-in; ok {
		if startTime, err = strconv.Atoi(s); err != nil {
			return "", fmt.Errorf("couldn't parse %q as start time: %w", s, err)
		}
	} else {
		return "", fmt.Errorf("couldn't read start time")
	}
	if s, ok := <-in; ok {
		for _, idstr := range strings.Split(s, ",") {
			if idstr == "x" {
				continue
			}
			idint, err := strconv.Atoi(idstr)
			if err != nil {
				return "", fmt.Errorf("couldn't parse %q as bus id: %w", idstr, err)
			}
			busIDs = append(busIDs, idint)
		}
	} else {
		return "", fmt.Errorf("couldn't read bus IDs")
	}
	var soonestBusID int
	nextBusTime := -1
	for j := range busIDs {
		next := (startTime - (startTime % busIDs[j])) + busIDs[j]
		if nextBusTime == -1 || next < nextBusTime {
			nextBusTime = next
			soonestBusID = busIDs[j]
		}
	}
	if nextBusTime == -1 {
		return "", fmt.Errorf("couldn't find next bus")
	}
	return strconv.Itoa((nextBusTime - startTime) * soonestBusID), nil
}

func part2func(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 13, part 1 - product of earliest bus ID and minutes to wait")
	ans, err := part1func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 13, part 2 - summary")
	ans, err = part2func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
