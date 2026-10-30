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

type policy struct {
	min int
	max int
	val rune
}

func parseInput(r io.Reader) ([]policy, []string, error) {
	var policies []policy
	var passwords []string

	scanner := bufio.NewScanner(r)
	line := 0

	for scanner.Scan() {
		line++
		fields := strings.Fields(scanner.Text())
		fmt.Printf("fields = %v\n", fields)
		if len(fields) != 3 {
			return nil, nil, fmt.Errorf("line %d: invalid input row\n", line)
		}
		// parse for policy
		minStr, maxStr, ok := strings.Cut(fields[0], "-")
		if !ok {
			return nil, nil, fmt.Errorf("line %d: invalid policy", line)
		}

		min, err := strconv.Atoi(minStr)
		if err != nil {
			return nil, nil, err
		}
		max, err := strconv.Atoi(maxStr)
		if err != nil {
			return nil, nil, err
		}
		// parse for value
		val := fields[1][0]

		pol := policy{min: min, max: max, val: rune(val)}
		policies = append(policies, pol)

		// parse for password
		passwords = append(passwords, fields[2])
	}

	return policies, passwords, scanner.Err()
}

func validPasswords(pol []policy, pass []string) int {
	var passwordCount int

	for i, p := range pol {
		valCount := 0
		for _, char := range pass[i] {
			if char == p.val {
				valCount++
			}
		}
		if valCount >= p.min && valCount <= p.max {
			passwordCount++
		}
	}
	return passwordCount
}

func validPasswords2(pol []policy, pass []string) int {
	var passwordCount int

	for i, p := range pol {
		first := rune(pass[i][p.min-1]) == p.val
		second := rune(pass[i][p.max-1]) == p.val

		if first != second {
			passwordCount++
		}
	}
	return passwordCount
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	policies, passwords, err := parseInput(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("valid passwords = %d\n", validPasswords(policies, passwords))
	fmt.Printf("valid passwords, part 2 = %d\n", validPasswords2(policies, passwords))
}
