package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

var regexMovement = regexp.MustCompile(`^(?P<direction>[A-Z])(?P<amount>[0-9]+)$`)

type Facing int

const (
	East  Facing = 0
	South Facing = 90
	West  Facing = 180
	North Facing = 270
)

type Position struct {
	X, Y int
	F    Facing
}

func (p *Position) eval(s string) error {
	r := regexMovement.FindStringSubmatch(s)
	if r == nil {
		return fmt.Errorf("cannot parse %q with %q", s, r)
	}
	i, err := strconv.Atoi(r[2])
	if err != nil {
		return fmt.Errorf("cannot parse integer part of %q: %w", s, err)
	}
	switch r[1] {
	case "L":
		if i%90 != 0 {
			return fmt.Errorf("cannot parse %q with non-whole left rotation %d", s, i)
		}
		p.F -= Facing(i)
		for p.F < 0 {
			p.F += Facing(360)
		}
	case "R":
		if i%90 != 0 {
			return fmt.Errorf("cannot parse %q with non-whole right rotation %d", s, i)
		}
		p.F += Facing(i)
		for p.F >= 360 {
			p.F -= Facing(360)
		}

	case "F":
		switch p.F {
		case North:
			r[1] = "N"
		case East:
			r[1] = "E"
		case South:
			r[1] = "S"
		case West:
			r[1] = "W"
		}
	}
	switch r[1] {
	case "N":
		p.Y -= i
	case "E":
		p.X += i
	case "S":
		p.Y += i
	case "W":
		p.X -= i
	}
	return nil
}

func part1func(in chan string) (string, error) {
	var err error
	p := &Position{}
	for s := range in {
		err = p.eval(s)
		if err != nil {
			return "", err
		}
	}
	var d int
	if p.X < 0 {
		d -= p.X
	} else {
		d += p.X
	}
	if p.Y < 0 {
		d -= p.Y
	} else {
		d += p.Y
	}

	return strconv.Itoa(d), nil
}

func part2func(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 12, part 1 - manhattan distance after movement")
	ans, err := part1func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 12, part 2 - manhattan distance after waypoint movement")
	ans, err = part2func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
