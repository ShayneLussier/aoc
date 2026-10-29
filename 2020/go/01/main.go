// https://adventofcode.com/2020/day/1

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(r io.Reader) ([]int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(string(data))
	nums := []int{}

	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func twoSum(expenses []int, target int) int {
	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			if expenses[j] == target-expenses[i] {
				return expenses[i] * expenses[j]
			}
		}
	}
	return 0
}

func threeSum(expenses []int, target int) int {
	for i := 0; i < len(expenses); i++ {
		if sum := twoSum(expenses[i+1:], target-expenses[i]); sum != 0 {
			return expenses[i] * sum
		}
	}
	return 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := parseInput(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part 1: %d\n", twoSum(data, 2020))   // 326211
	fmt.Printf("part 2: %d\n", threeSum(data, 2020)) // 131347190
}
