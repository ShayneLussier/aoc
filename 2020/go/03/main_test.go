package main

import (
	"os"
	"slices"
	"strings"
	"testing"
)

const sampleInput = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestParseInput(t *testing.T) {
	got := parseInput(strings.NewReader(sampleInput))
	want := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	if !slices.Equal(got, want) {
		t.Errorf("got = %v, want = %v\n", got, want)
	}
}

func TestTreesEncountered(t *testing.T) {
	sample := parseInput(strings.NewReader(sampleInput))
	f, err := os.Open("input.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	actual := parseInput(f)

	tests := []struct {
		name  string
		slope []string
		angle angle
		want  int
	}{
		{name: "sample", slope: sample, angle: angle{right: 3, down: 1}, want: 7},
		{name: "actual", slope: actual, angle: angle{right: 3, down: 1}, want: 286},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := treesEncountered(tt.slope, tt.angle)

			if got != tt.want {
				t.Errorf("got = %d, want = %d\n", got, tt.want)
			}
		})
	}
}

func TestTreesMult(t *testing.T) {
	sample := parseInput(strings.NewReader(sampleInput))

	angles := []angle{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	got := treesMult(sample, angles)
	want := 336

	if got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}
