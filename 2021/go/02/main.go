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

type move struct {
	dir string
	n   int
}

func parseMoves(r io.Reader) ([]move, error) {
	var moves []move
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			return nil, fmt.Errorf("bad line: %q", scanner.Text())
		}
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		moves = append(moves, move{fields[0], n})
	}
	return moves, scanner.Err()
}

func calcDepth(moves []move) int {
	var horizontal, depth int

	for _, m := range moves {
		switch m.dir {
		case "forward":
			horizontal += m.n
		case "up":
			depth -= m.n
		case "down":
			depth += m.n
		}
	}

	return horizontal * depth
}
func calcAim(moves []move) int {
	var horizontal, depth, aim int

	for _, m := range moves {
		switch m.dir {
		case "forward":
			horizontal += m.n
			depth += aim * m.n
		case "up":
			aim -= m.n
		case "down":
			aim += m.n
		}
	}

	return horizontal * depth
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	moves, err := parseMoves(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part 1 = %v\n", calcDepth(moves))
	fmt.Printf("part 2 = %v\n", calcAim(moves))
}
