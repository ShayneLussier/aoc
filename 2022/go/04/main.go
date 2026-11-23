// https://adventofcode.com/2022/day/4

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type section struct {
	lo, hi int
}

type pair [2]section

func parseInput(r io.Reader) ([]pair, error) {
	var pairs []pair

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var p pair
		_, err := fmt.Sscanf(strings.TrimSpace(scanner.Text()), "%d-%d,%d-%d", &p[0].lo, &p[0].hi, &p[1].lo, &p[1].hi)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, p)
	}
	return pairs, scanner.Err()
}

func parseSections(pairs []pair) (overlaps int) {
	for _, pair := range pairs {
		if pair[0].lo <= pair[1].lo && pair[0].hi >= pair[1].hi {
			overlaps++
			continue
		}
		if pair[1].lo <= pair[0].lo && pair[1].hi >= pair[0].hi {
			overlaps++
			continue
		}
	}

	return overlaps
}

func parseSections2(pairs []pair) (overlaps int) {
	for _, p := range pairs {
		if p[0].lo <= p[1].hi && p[1].lo <= p[0].hi {
			overlaps++
		}
	}
	return overlaps
}

func main() {
	// In how many assignment pairs does one range fully contain the other?
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pairs, err := parseInput(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There are %d section overlaps.\n", parseSections(pairs)) // 569
	fmt.Printf("There are %d range overlaps.\n", parseSections2(pairs))  // 936
}
