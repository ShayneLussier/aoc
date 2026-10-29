// https://adventofcode.com/2021/day/3

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

func parseInput(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func countBits(lines []string) ([]int, []int) {
	ones := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				ones[i]++
			}
		}
	}
	zeros := make([]int, len(lines[0]))
	for i := range zeros {
		zeros[i] = len(lines) - ones[i]
	}
	return ones, zeros
}

func powerConsumption(ones, zeros []int) (int, error) {
	var gamma, epsilon strings.Builder

	for i := range len(ones) {
		if ones[i] > zeros[i] {
			gamma.WriteByte('1')
			epsilon.WriteByte('0')
		} else {
			gamma.WriteByte('0')
			epsilon.WriteByte('1')
		}
	}
	g, err := strconv.ParseInt(gamma.String(), 2, 64)
	if err != nil {
		return 0, err
	}
	e, err := strconv.ParseInt(epsilon.String(), 2, 64)
	if err != nil {
		return 0, err
	}
	return int(g * e), nil
}

func lifeSupport(lines []string, wantedBit rune) (int, error) {
	for i := 0; len(lines) > 1; i++ {
		ones, zeros := countBits(lines)

		var want byte

		switch wantedBit {
		case '1':
			if ones[i] >= zeros[i] {
				want = '1'
			} else {
				want = '0'
			}
		case '0':
			if ones[i] >= zeros[i] {
				want = '0'
			} else {
				want = '1'
			}
		}

		var kept []string
		for _, line := range lines {
			if line[i] == want {
				kept = append(kept, line)
			}
		}
		lines = kept
	}
	// parse the last line
	n, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	diag, err := parseInput(f)
	if err != nil {
		log.Fatal(err)
	}

	power, err := powerConsumption(countBits(diag))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1 = %d\n", power)

	oxygen, err := lifeSupport(diag, '1')
	if err != nil {
		log.Fatal(err)
	}
	CO2, err := lifeSupport(diag, '0')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part 2 = %d\n", oxygen*CO2)
}
