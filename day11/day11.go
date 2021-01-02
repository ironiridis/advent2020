package main

import (
	"fmt"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

type seat rune
type seatMap [][]seat
type mapUpdate struct {
	row, col int
	set      seat
}

const (
	Occupied seat = '#'
	Empty    seat = 'L'
	Floor    seat = '.'
	Wall     seat = '!'
)

func newSeatMap(in chan string) seatMap {
	var r seatMap
	for s := range in {
		if s == "" {
			continue
		}
		r = append(r, []seat(s))
	}
	return r
}

func (sm seatMap) at(r, c int) seat {
	if r < 0 || c < 0 || r >= len(sm) || c >= len(sm[r]) {
		return Wall
	}
	return sm[r][c]
}

func (sm seatMap) occupiedAround(r, c int) (o int) {
	if sm.at(r-1, c-1) == Occupied {
		o++
	}
	if sm.at(r-1, c) == Occupied {
		o++
	}
	if sm.at(r-1, c+1) == Occupied {
		o++
	}

	if sm.at(r, c-1) == Occupied {
		o++
	}
	if sm.at(r, c+1) == Occupied {
		o++
	}

	if sm.at(r+1, c-1) == Occupied {
		o++
	}
	if sm.at(r+1, c) == Occupied {
		o++
	}
	if sm.at(r+1, c+1) == Occupied {
		o++
	}
	return
}

func (sm seatMap) getPart1Updates() []mapUpdate {
	var u []mapUpdate
	for r := range sm {
		for c := range sm[r] {
			switch sm.at(r, c) {
			case Occupied:
				if sm.occupiedAround(r, c) >= 4 {
					u = append(u, mapUpdate{row: r, col: c, set: Empty})
				}
			case Empty:
				if sm.occupiedAround(r, c) == 0 {
					u = append(u, mapUpdate{row: r, col: c, set: Occupied})
				}
			}
		}
	}
	return u
}

func (sm seatMap) applyUpdates(u []mapUpdate) {
	for idx := range u {
		sm[u[idx].row][u[idx].col] = u[idx].set
	}
}

func (sm seatMap) countOccupied() (o int) {
	for r := range sm {
		for c := range sm[r] {
			if sm.at(r, c) == Occupied {
				o++
			}
		}
	}
	return

}

func part1func(in chan string) (string, error) {
	sm := newSeatMap(in)
	for {
		upd := sm.getPart1Updates()
		if upd == nil {
			break
		}
		sm.applyUpdates(upd)
	}

	return strconv.Itoa(sm.countOccupied()), nil
}

func part2func(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day 11, part 1 - occupied seats after settled")
	ans, err := part1func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 11, part 2 - occupied seats after settled, but more complicated")
	ans, err = part2func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
