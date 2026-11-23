package main

import (
	"slices"
	"testing"
)

const example = `
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestParseInput(t *testing.T) {
	stacks, instructions := parseInput(example)

	wantStacks := []stack{{"Z", "N"}, {"M", "C", "D"}, {"P"}}
	if !slices.EqualFunc(stacks, wantStacks, slices.Equal) {
		t.Errorf("stacks = %v, want = %v", stacks, wantStacks)
	}

	wantInstructions := []instruction{{1, 2, 1}, {3, 1, 3}, {2, 2, 1}, {1, 1, 2}}
	if !slices.Equal(instructions, wantInstructions) {
		t.Errorf("instructions = %v, want = %v", instructions, wantInstructions)
	}
}

func TestCrateMovers(t *testing.T) {
	tests := []struct {
		name  string
		mover func([]stack, []instruction) string
		want  string
	}{
		{"crateMover9000", crateMover9000, "CMZ"},
		{"crateMover9001", crateMover9001, "MCD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stacks, instructions := parseInput(example) // to get unmutated inputs
			if got := tt.mover(stacks, instructions); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
