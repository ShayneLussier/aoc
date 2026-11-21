// https://adventofcode.com/2022/day/1

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

func parseElves(r io.Reader) (elves []int, err error) {
	var current int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elves = append(elves, current)
			current = 0
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		current += cal
	}
	elves = append(elves, current) // handle last elf

	slices.Sort(elves)
	return elves, scanner.Err()
}

func sumTop(elves []int, n int) int {
	var total int
	for _, cal := range elves[len(elves)-n:] {
		total += cal
	}
	return total
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	elves, err := parseElves(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The top elf has %d calories\n", sumTop(elves, 1))      // 72478
	fmt.Printf("The top 3 elves have %d calories\n", sumTop(elves, 3)) // 210367
}
