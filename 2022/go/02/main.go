// https://adventofcode.com/2022/day/2

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseGuide(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

func scorePart1(j, k byte) int {
	switch {
	case j == 'A' && k == 'X':
		return 4
	case j == 'A' && k == 'Y':
		return 8
	case j == 'A' && k == 'Z':
		return 3
	case j == 'B' && k == 'X':
		return 1
	case j == 'B' && k == 'Y':
		return 5
	case j == 'B' && k == 'Z':
		return 9
	case j == 'C' && k == 'X':
		return 7
	case j == 'C' && k == 'Y':
		return 2
	case j == 'C' && k == 'Z':
		return 6
	}
	return 0
}

func scorePart2(j, k byte) int {
	switch {
	case j == 'A' && k == 'X':
		return 3
	case j == 'A' && k == 'Y':
		return 4
	case j == 'A' && k == 'Z':
		return 8
	case j == 'B' && k == 'X':
		return 1
	case j == 'B' && k == 'Y':
		return 5
	case j == 'B' && k == 'Z':
		return 9
	case j == 'C' && k == 'X':
		return 2
	case j == 'C' && k == 'Y':
		return 6
	case j == 'C' && k == 'Z':
		return 7
	}
	return 0
}

func calculatePoints(guide []string) (pts1, pts2 int) {
	for _, line := range guide {
		opponentMove := line[0]
		strategy := line[len(line)-1]

		pts1 += scorePart1(opponentMove, strategy)
		pts2 += scorePart2(opponentMove, strategy)
	}
	return pts1, pts2
}

func main() {
	guide, err := parseGuide("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	p1, p2 := calculatePoints(guide)

	fmt.Printf("By playing perfectly, the player get's %d points!\n", p1)         // 14264
	fmt.Printf("By knowing the true strategy, the player earns %d points!\n", p2) // 12382
}
