package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

type instruction struct {
	sym string
	val int
}

type programState struct {
	pc           int
	acc          int
	instructions []*instruction
}

var regexParseInstruction = regexp.MustCompile(`^(?P<symbol>[a-z]{3}) \+?(?P<value>-?[0-9]+)$`)

func parseInstruction(t string) (*instruction, error) {
	r := regexParseInstruction.FindStringSubmatch(t)
	if r == nil {
		return nil, fmt.Errorf("cannot parse instruction %q using %q", t, r)
	}
	v, err := strconv.Atoi(r[2])
	if err != nil {
		return nil, fmt.Errorf("cannot parse value of instruction %q: %w", t, err)
	}
	return &instruction{sym: r[1], val: v}, nil
}

func newProgramState(in chan string) (*programState, error) {
	ps := &programState{}
	for t := range in {
		if t == "" {
			continue
		}
		ins, err := parseInstruction(t)
		if err != nil {
			return nil, fmt.Errorf("cannot create programState: %w", err)
		}
		ps.instructions = append(ps.instructions, ins)
	}
	return ps, nil
}

func (ps *programState) eval() (int, error) {
	if ps.instructions[ps.pc] == nil {
		return ps.acc, fmt.Errorf("no instruction at pc=%d", ps.pc)
	}
	switch ps.instructions[ps.pc].sym {
	case "nop":
		ps.pc++
	case "acc":
		ps.acc += ps.instructions[ps.pc].val
		ps.pc++
	case "jmp":
		ps.pc += ps.instructions[ps.pc].val
	}
	return ps.acc, nil
}

func part1func(in chan string) (string, error) {
	ps, err := newProgramState(in)
	if err != nil {
		return "", err
	}
	var steps int
	visited := make([]bool, len(ps.instructions))
	for {
		steps++
		visited[ps.pc] = true
		acc, err := ps.eval()
		if err != nil {
			return "", fmt.Errorf("couldn't evaluate instruction: %w", err)
		}
		if steps > len(ps.instructions) {
			return "", fmt.Errorf("didn't terminate within expected number of steps")
		}
		if visited[ps.pc] {
			return strconv.Itoa(acc), nil
		}
	}
}

func part2func(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 8, part 1 - accumulator value at repeat")
	ans, err := part1func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 8, part 2 - summary")
	ans, err = part2func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
