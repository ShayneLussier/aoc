// https://adventofcode.com/2020/day/4

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid *string // optional, can be nil
}

func validYear(s string, min, max int) bool {
	n, err := strconv.Atoi(s)
	return len(s) == 4 && err == nil && n >= min && n <= max
}

func validHeight(s string) bool {
	if len(s) < 3 { // missing field or no unit
		return false
	}
	height, err := strconv.Atoi(s[:len(s)-2])
	if err != nil {
		return false
	}
	unit := s[len(s)-2:]

	switch unit {
	case "cm":
		if height >= 150 && height <= 193 {
			return true
		}
	case "in":
		if height >= 59 && height <= 76 {
			return true
		}
	}
	return false
}

func validHairColor(s string) bool {
	if len(s) != 7 {
		return false
	}
	if s[0] != '#' {
		return false
	}

	for _, b := range s[1:] {
		if !(b >= '0' && b <= '9' || b >= 'a' && b <= 'f') {
			return false
		}
	}
	return true
}

func validEyeColor(s string) bool {
	switch s {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func validPID(s string) bool {
	if len(s) != 9 {
		return false
	}
	for _, b := range s {
		if b < '0' || b > '9' {
			return false
		}
	}
	return true
}

func (p *passport) hasFields() bool {
	return p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func (p *passport) validate() bool {
	return validYear(p.byr, 1920, 2002) &&
		validYear(p.iyr, 2010, 2020) &&
		validYear(p.eyr, 2020, 2030) &&
		validHeight(p.hgt) &&
		validHairColor(p.hcl) &&
		validEyeColor(p.ecl) &&
		validPID(p.pid)
}

func parseInput(r io.Reader) ([]passport, error) {
	var passports []passport

	p := passport{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if scanner.Text() == "" {
			passports = append(passports, p)
			p = passport{}
		}

		for f := range strings.FieldsSeq(scanner.Text()) {
			key, value, found := strings.Cut(f, ":")

			if !found {
				continue
			}

			switch key {
			case "byr":
				p.byr = value
			case "iyr":
				p.iyr = value
			case "eyr":
				p.eyr = value
			case "hgt":
				p.hgt = value
			case "hcl":
				p.hcl = value
			case "ecl":
				p.ecl = value
			case "pid":
				p.pid = value
			case "cid":
				p.cid = &value
			}
		}
	}
	// append the last
	passports = append(passports, p)

	return passports, scanner.Err()
}

func parsePassports(pp []passport) (part1, part2 int) {
	for _, p := range pp {
		if p.hasFields() {
			part1++
		}
		if p.validate() {
			part2++
		}
	}
	return part1, part2
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	passports, err := parseInput(f)
	if err != nil {
		log.Fatal(err)
	}

	part1, part2 := parsePassports(passports)
	fmt.Printf("Part 1: %d\n", part1) // 204
	fmt.Printf("Part 2: %d\n", part2) // 179
}
