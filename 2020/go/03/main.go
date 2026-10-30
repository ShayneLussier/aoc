// https://adventofcode.com/2020/day/3

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type angle struct {
	right, down int
}

func parseInput(r io.Reader) []string {
	var slope []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		slope = append(slope, scanner.Text())
	}

	return slope
}

func treesEncountered(slope []string, a angle) int {
	var trees, x int

	for i := 0; i < len(slope); i += a.down {
		// wrap around if out of bounds
		x %= len(slope[i])
		// check if there is a tree '#'
		if slope[i][x] == '#' {
			trees++
		}
		x += a.right
	}
	return trees
}

func treesMult(slope []string, angles []angle) int {
	result := 1

	for _, a := range angles {
		result *= treesEncountered(slope, a)
	}

	return result
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	slope := parseInput(f)

	fmt.Printf("trees encountered = %d\n", treesEncountered(slope, angle{right: 3, down: 1})) // 286

	angles := []angle{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	fmt.Printf("mulitplied trees encountered = %d\n", treesMult(slope, angles)) // 3638606400
}
