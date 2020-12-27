package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ironiridis/advent2020/scando"
)

var validateHeight = regexp.MustCompile(`^([0-9]{2,3})(in|cm)$`)
var validateHair = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var validateEyes = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
var validatePID = regexp.MustCompile(`^[0-9]{9}$`)

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

func (p passport) validFieldsPresent() bool {
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

func (p passport) validValues() bool {
	if byr, err := strconv.Atoi(p["byr"]); err != nil || byr < 1920 || byr > 2002 {
		return false
	}

	if iyr, err := strconv.Atoi(p["iyr"]); err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}

	if eyr, err := strconv.Atoi(p["eyr"]); err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}

	hgt := validateHeight.FindStringSubmatch(p["hgt"])
	if hgt == nil {
		return false
	}
	hgtn, err := strconv.Atoi(hgt[1])
	if err != nil {
		return false
	}
	switch hgt[2] {
	case "cm":
		if hgtn < 150 || hgtn > 193 {
			return false
		}
	case "in":
		if hgtn < 59 || hgtn > 76 {
			return false
		}
	default: /* unrecognized unit, though should be covered by regex */
		return false
	}
	if !validateHair.MatchString(p["hcl"]) {
		return false
	}
	if !validateEyes.MatchString(p["ecl"]) {
		return false
	}
	if !validatePID.MatchString(p["pid"]) {
		return false
	}
	return true
}

func (p passport) valid() bool {
	return p.validFieldsPresent() && p.validValues()
}

func countValidPassportsPart1(in chan string) (string, error) {
	var valid int
	for {
		p, done, err := readPassport(in)
		if err != nil {
			return "", err
		}
		if p.validFieldsPresent() {
			valid++
		}
		if done {
			break
		}
	}
	return strconv.Itoa(valid), nil
}

func countValidPassportsPart2(in chan string) (string, error) {
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

func main() {
	fmt.Println("Day 4, part 1 - valid passport count, just checking field presence")
	ans, err := countValidPassportsPart1(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day 4, part 2 - valid passport count, with value validation")
	ans, err = countValidPassportsPart2(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
