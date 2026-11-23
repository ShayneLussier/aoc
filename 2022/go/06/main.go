// https://adventofcode.com/2022/day/6

package main

import (
	"fmt"
	"log"
	"os"
)

func parseStream(stream string, n int) int {
	for i := 0; i < len(stream); i++ {
		seen := map[rune]struct{}{}
		for _, c := range stream[i : i+n] {
			seen[c] = struct{}{}
		}
		// if seen has n distinct characters, return the last's index
		if len(seen) == n {
			return i + n
		}
	}
	return 0
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d characters were processed before the start-of-packet marker was detected.\n", parseStream(string(data), 4))   // 1343
	fmt.Printf("%d characters were processed before the start-of-message marker was detected.\n", parseStream(string(data), 14)) // 2193
}
