// https://adventofcode.com/2019/day/1

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func parseInput(r io.Reader) ([]int, error) {
	var masses []int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return masses, err
		}
		masses = append(masses, mass)
	}
	return masses, scanner.Err()
}

func fuelSum(mm []int) (simple, recursive int) {
	for _, m := range mm {
		simple += m/3 - 2
		recursive += fuelRequirement(m)
	}
	return simple, recursive
}

func fuelRequirement(m int) int {
	f := m/3 - 2
	if f <= 0 {
		return 0
	}
	return f + fuelRequirement(f)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	masses, err := parseInput(f)
	if err != nil {
		log.Fatal(err)
	}

	simple, recursive := fuelSum(masses)
	fmt.Printf("The simple sum of fuel required is %d\n", simple)
	fmt.Printf("The recursive sum of fuel required is %d\n", recursive)

}
