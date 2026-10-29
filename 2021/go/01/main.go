// https://adventofcode.com/2021/day/1

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func readInts(r io.Reader) ([]int, error) {
	var nums []int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, scanner.Err()
}

func countIncreases(nums []int, window int) (int, error) {
	if window < 1 || len(nums) < window {
		return 0, fmt.Errorf("window %d invalid for %d readings", window, len(nums))
	}

	var increases int

	for i := range len(nums) - window {
		if nums[i+window] > nums[i] {
			increases++
		}
	}

	return increases, nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	nums, err := readInts(f)
	if err != nil {
		log.Fatal(err)
	}

	part1, err := countIncreases(nums, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1 = %v\n", part1)

	part2, err := countIncreases(nums, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 2 = %v\n", part2)
}
