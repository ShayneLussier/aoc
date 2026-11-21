// https://adventofcode.com/2019/day/2

package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(s string) ([]int, error) {
	var codes []int
	for s := range strings.SplitSeq(strings.TrimSpace(s), ",") {
		c, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		codes = append(codes, c)
	}
	return codes, nil
}

func parseCodes(codes []int) int {
	for i := 0; i < len(codes); i += 4 {
		switch codes[i] {
		case 1:
			codes[codes[i+3]] = codes[codes[i+1]] + codes[codes[i+2]]
		case 2:
			codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
		case 99:
			return codes[0]
		default:
			panic(fmt.Sprintf("unknown opcode = %d", codes[i]))
		}
	}
	return 0
}

func determineInputPair(codes []int) (int, error) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			part2 := slices.Clone(codes)
			part2[1], part2[2] = noun, verb
			if parseCodes(part2) == 19690720 {
				return 100*noun + verb, nil
			}
		}
	}
	return 0, fmt.Errorf("not input pairs produce the correct output")
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	codes, err := parseInput(string(data))
	if err != nil {
		log.Fatal(err)
	}

	part1 := slices.Clone(codes)
	part1[1], part1[2] = 12, 2                                       // "1202 program alarm" state
	fmt.Printf("The intial value left is = %d\n", parseCodes(part1)) // 3101878

	part2, err := determineInputPair(slices.Clone(codes))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The inputs who produce an output of 19690720 return an answer of = %d\n", part2) // 8444
}
