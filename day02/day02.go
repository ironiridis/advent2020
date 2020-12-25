package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ironiridis/advent2020/scando"
)

func countValidPasswords(in chan string) (string, error) {
	infmt := regexp.MustCompile(`^(?P<minimum>[0-9]+)-(?P<maximum>[0-9]+) (?P<symbol>.): (?P<subject>.+)$`)
	valid := 0
	for s := range in {
		parse := infmt.FindStringSubmatch(s)
		if parse == nil {
			return "", fmt.Errorf("unable to parse line %q with %q", s, infmt)
		}
		min, _ := strconv.Atoi(parse[1])
		max, _ := strconv.Atoi(parse[2])
		symInSubject := strings.Count(parse[4], parse[3])
		if symInSubject >= min && symInSubject <= max {
			valid++
		}
	}
	return strconv.Itoa(valid), nil
}

func alternateCountValidPasswords(in chan string) (string, error) {
	infmt := regexp.MustCompile(`^(?P<pos1>[0-9]+)-(?P<pos2>[0-9]+) (?P<symbol>.): (?P<subject>.+)$`)
	valid := 0
	for s := range in {
		parse := infmt.FindStringSubmatch(s)
		if parse == nil {
			return "", fmt.Errorf("unable to parse line %q with %q", s, infmt)
		}
	}
	return strconv.Itoa(valid), nil
}

func main() {
	fmt.Println("Day 2, part 1 - password validity, symbol count ranges")
	ans, err := countValidPasswords(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 2, part 2 - password validity, mutually exclusive symbol positions")
	ans, err := alternateCountValidPasswords(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
