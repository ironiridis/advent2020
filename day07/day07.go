package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

var regexRule = regexp.MustCompile(`^(?P<container>\S+ \S+) bags contain (?P<contents>.+)\.$`)
var regexBag = regexp.MustCompile(`(?P<count>[0-9]+) (?P<container>\S+ \S+) bags?`)

type contain map[string]int
type rules map[string]contain

func (r rules) parse(s string) (err error) {
	if s == "" {
		return
	}
	ln := regexRule.FindStringSubmatch(s)
	if ln == nil {
		err = fmt.Errorf("cannot parse %q with %q", s, regexRule)
		return
	}
	cont := regexBag.FindAllStringSubmatch(s, -1)

	if r[ln[1]] != nil {
		err = fmt.Errorf("redefining rule for %q", ln[1])
		return
	}
	r[ln[1]] = make(contain, len(cont))
	for k := range cont {
		v, _ := strconv.Atoi(cont[k][1]) // regex implies [0-9]+ here, ignore err
		r[ln[1]][cont[k][2]] = v
	}
	return
}

func (r rules) parseAll(in chan string) (err error) {
	for s := range in {
		err = r.parse(s)
		if err != nil {
			break
		}
	}
	return
}

func (r rules) deleteBag(b string) {
	delete(r, b)       // delete the direct reference
	for o := range r { // scan through all container rules
		delete(r[o], b)     // delete the reference to bag b in each rule
		if len(r[o]) == 0 { // if the rule is now empty
			r.deleteBag(o) // recurse
		}
	}
}

func part1func(in chan string) (string, error) {
	r := make(rules)
	err := r.parseAll(in)
	if err != nil {
		return "", err
	}

	delete(r, "shiny gold") // target color; remove it from container map
	for o := range r {
		if len(r[o]) == 0 { // if the rule for this bag color is empty
			r.deleteBag(o) // delete this bag color from the map
		}
	}

	return strconv.Itoa(len(r)), nil
}

func part2func(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 7, part 1 - summary")
	ans, err := part1func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 7, part 2 - summary")
	ans, err = part2func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
