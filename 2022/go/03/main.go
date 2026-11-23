// https://adventofcode.com/2022/day/3

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func intersection[T comparable](slices ...[]T) (T, bool) {
	var zero T

	if len(slices) == 0 {
		return zero, false
	}

	counts := make(map[T]int)

	for _, s := range slices {
		seen := make(map[T]struct{})

		for _, c := range s {
			_, exists := seen[c]
			if !exists {
				seen[c] = struct{}{}
				counts[c]++
			}
		}
	}

	for _, c := range slices[0] {
		if counts[c] == len(slices) {
			return c, true
		}
	}

	return zero, false
}

func parseInput(ruckacks []string) (priority int, err error) {
	for _, s := range ruckacks {
		l := len(s) / 2
		c1 := s[:l]
		c2 := s[l:]

		item, ok := intersection([]byte(c1), []byte(c2))
		if !ok {
			return 0, fmt.Errorf("no common item between all rucksacks")
		}

		switch {
		case 'a' <= item && item <= 'z':
			priority += int(item) % 96
		case 'A' <= item && item <= 'Z':
			priority += (int(item) % 65) + 27
		}
	}
	return priority, nil
}

func parseBadges(rucksacks []string) (priority int, err error) {
	i := 0
	for i <= len(rucksacks)-3 {
		c1 := rucksacks[i]
		c2 := rucksacks[i+1]
		c3 := rucksacks[i+2]

		item, ok := intersection([]byte(c1), []byte(c2), []byte(c3))
		if !ok {
			return 0, fmt.Errorf("no common item between all rucksacks")
		}

		switch {
		case 'a' <= item && item <= 'z':
			priority += int(item) % 96
		case 'A' <= item && item <= 'Z':
			priority += (int(item) % 65) + 27
		}
		i += 3
	}
	return priority, nil
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rucksacks := strings.Fields(string(f))

	sum1, err := parseInput(rucksacks)
	if err != nil {
		log.Fatal(err)
	}
	sum2, err := parseBadges(rucksacks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 = %d.\n", sum1) // 8139
	fmt.Printf("Part 2 = %d.\n", sum2) // 2668
}
