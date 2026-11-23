// https://adventofcode.com/2022/day/5

package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type instruction struct {
	qty, src, dst int
}

type stack = []string

func parseStacks(data string) (stacks []stack) {
	d := strings.Split(data, "\n")

	// prepare n amount of empty stacks
	stackCount := strings.Fields(d[len(d)-1])
	for range stackCount {
		stacks = append(stacks, stack{})
	}

	for _, line := range d {
		i := 0
		j := 0
		for j < len(line) {
			symbol := line[j+1]
			if symbol >= 'A' && symbol <= 'Z' {
				stacks[i] = append(stacks[i], string(symbol))
			}
			i++
			j += 4
		}
	}

	for _, s := range stacks {
		slices.Reverse(s)
	}

	return stacks
}

func parseInstructions(data string) (instructions []instruction) {
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		var i instruction
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &i.qty, &i.src, &i.dst)
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, i)
	}
	return instructions
}

func parseInput(data string) ([]stack, []instruction) {
	d := strings.Split(data, "\n\n")

	stacks := parseStacks(d[0])
	instructions := parseInstructions(d[1])

	return stacks, instructions
}

func crateMover9000(stacks []stack, instructions []instruction) (topCrates string) {
	for _, i := range instructions {
		src := i.src - 1
		dst := i.dst - 1

		for range i.qty {
			crate := stacks[src][len(stacks[src])-1]       // retrieve the top crate
			stacks[src] = stacks[src][:len(stacks[src])-1] // pop it from the slice
			stacks[dst] = append(stacks[dst], crate)       // append to new stack
		}
	}

	for _, s := range stacks {
		if len(s) > 0 {
			topCrates += s[len(s)-1]
		}
	}

	return topCrates
}

func crateMover9001(stacks []stack, instructions []instruction) (topCrates string) {
	for _, i := range instructions {
		src := i.src - 1
		dst := i.dst - 1

		var crane stack
		for range i.qty {
			crate := stacks[src][len(stacks[src])-1]       // retrieve the top crate
			stacks[src] = stacks[src][:len(stacks[src])-1] // pop it from the slice
			crane = append(crane, crate)                   // append to new crane
		}

		slices.Reverse(crane)
		stacks[dst] = append(stacks[dst], crane...)
	}

	for _, s := range stacks {
		if len(s) > 0 {
			topCrates += s[len(s)-1]
		}
	}

	return topCrates
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	stacks, instructions := parseInput(string(data))
	stacksClone, _ := parseInput(string(data)) // parse Input a second to avoir a shallow slice clone

	fmt.Printf("The crate of the top of every stack makes code = %v.\n", crateMover9000(stacks, instructions))      // GFTNRBZPF
	fmt.Printf("The crate of the top of every stack makes code = %v.\n", crateMover9001(stacksClone, instructions)) // VRQWPDSGP
}
