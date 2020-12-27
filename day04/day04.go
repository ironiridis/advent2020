package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

type passport map[string]string

func readPassport(in chan string) (p passport, done bool, err error) {
	infmt, err := regexp.Compile(`(?P<key>[a-z]{3}):(?P<val>\S+)`)
	if err != nil {
		return
	}

	for {
		ln, ok := <-in
		if !ok { // channel is closed, no more data
			done = true
		}
		if ln == "" { // passport is finished
			return
		}
		parsed := infmt.FindAllStringSubmatch(ln, -1)
		if parsed == nil {
			err = fmt.Errorf("unable to parse %q using %v", ln, infmt)
			return
		}
		for midx := range parsed {
			if p == nil {
				p = make(passport, 8)
			}
			p[parsed[midx][1]] = parsed[midx][2]
		}
	}
}

func (p passport) valid() bool {
	if p == nil {
		return false
	}
	if p["byr"] == "" {
		return false
	}
	if p["iyr"] == "" {
		return false
	}
	if p["eyr"] == "" {
		return false
	}
	if p["hgt"] == "" {
		return false
	}
	if p["hcl"] == "" {
		return false
	}
	if p["ecl"] == "" {
		return false
	}
	if p["pid"] == "" {
		return false
	}
	// cid not required
	return true
}

func countValidPassportsPart1(in chan string) (string, error) {
	var valid int
	for {
		p, done, err := readPassport(in)
		if err != nil {
			return "", err
		}
		if p.valid() {
			valid++
		}
		if done {
			break
		}
	}
	return strconv.Itoa(valid), nil
}

func countValidPassportsPart2(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 3, part 1 - valid passport count")
	ans, err := countValidPassportsPart1(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 3, part 2 ...")
	ans, err = countValidPassportsPart2(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
