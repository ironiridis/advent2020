package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

var regexMovement = regexp.MustCompile(`^(?P<action>[A-Z])(?P<amount>[0-9]+)$`)

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

type Waypoint struct {
	X, Y int
}

func (p *Position) evalDirect(s string) error {
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

func (wp *Waypoint) RotateCW() {
	wp.X, wp.Y = 0-wp.Y, wp.X
}
func (wp *Waypoint) RotateCCW() {
	wp.X, wp.Y = wp.Y, 0-wp.X
}

func (p *Position) evalWithWaypoint(s string, wp *Waypoint) error {
	r := regexMovement.FindStringSubmatch(s)
	if r == nil {
		return fmt.Errorf("cannot parse %q with %q", s, r)
	}
	i, err := strconv.Atoi(r[2])
	if err != nil {
		return fmt.Errorf("cannot parse integer part of %q: %w", s, err)
	}
	if i < 0 {
		return fmt.Errorf("this code doesn't handle negative values correctly, so we cannot parse %q", s)
	}
	switch r[1] {
	case "N":
		wp.Y -= i
	case "E":
		wp.X += i
	case "S":
		wp.Y += i
	case "W":
		wp.X -= i
	case "L":
		if i%90 != 0 {
			return fmt.Errorf("cannot parse %q with non-whole left rotation %d", s, i)
		}
		for i > 0 {
			wp.RotateCCW()
			i -= 90
		}
	case "R":
		if i%90 != 0 {
			return fmt.Errorf("cannot parse %q with non-whole right rotation %d", s, i)
		}
		for i > 0 {
			wp.RotateCW()
			i -= 90
		}
	case "F":
		for ; i > 0; i-- {
			p.X += wp.X
			p.Y += wp.Y
		}
	}
	return nil
}

func manhattan(x, y int) (d int) {
	if x < 0 {
		d -= x
	} else {
		d += x
	}
	if y < 0 {
		d -= y
	} else {
		d += y
	}
	return
}

func part1func(in chan string) (string, error) {
	var err error
	p := &Position{}
	for s := range in {
		err = p.evalDirect(s)
		if err != nil {
			return "", err
		}
	}
	return strconv.Itoa(manhattan(p.X, p.Y)), nil
}

func part2func(in chan string) (string, error) {
	var err error
	p := &Position{}
	wp := &Waypoint{X: 10, Y: -1}
	for s := range in {
		err = p.evalWithWaypoint(s, wp)
		if err != nil {
			return "", err
		}
	}
	return strconv.Itoa(manhattan(p.X, p.Y)), nil
}

func main() {
	fmt.Println("Day 12, part 1 - manhattan distance after direct movement")
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
