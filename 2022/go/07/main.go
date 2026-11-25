// https://adventofcode.com/2022/day/7

package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func dirSizes(c []string) map[string]int {
	sizes := make(map[string]int)
	var currentDir []string

	for _, line := range c {
		tokens := strings.Fields(string(line))
		if len(tokens) == 0 {
			continue
		}

		switch {
		case tokens[0] == "$" && tokens[1] == "cd" && tokens[2] == "..":
			currentDir = currentDir[:len(currentDir)-1]
		case tokens[0] == "$" && tokens[1] == "cd":
			currentDir = append(currentDir, tokens[2])
		default:
			if unicode.IsDigit(rune(tokens[0][0])) {
				var path string
				for i := range len(currentDir) {
					for _, p := range currentDir[:i+1] {
						path += p
					}
					s, err := strconv.Atoi(tokens[0])
					if err != nil {
						panic(err)
					}
					sizes[path] += s
					path = ""
				}
			}
		}
	}
	return sizes
}

func part1(sizes map[string]int) int {
	var total int

	for _, s := range sizes {
		if s < 100_000 {
			total += s
		}
	}
	return total
}

func part2(sizes map[string]int) int {
	needed := 30_000_000 - (70_000_000 - sizes["/"])
	var eligible []int

	for _, s := range sizes {
		if s > needed {
			eligible = append(eligible, s)
		}
	}
	return slices.Min(eligible)
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	commands := strings.Split(string(contents), "\n")

	fmt.Printf("Part 1 = %d!\n", part1(dirSizes(commands))) // 1325919
	fmt.Printf("Part 2 = %d!\n", part2(dirSizes(commands))) // 2050735
}
