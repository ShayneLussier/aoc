package main

import (
	"slices"
	"strings"
	"testing"
)

// Example from the puzzle description.
const example = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestParseElves(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{"example", example, []int{4000, 6000, 10000, 11000, 24000}, false},
		{"single elf", "5000", []int{5000}, false},
		{"Atoi error", "hello", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseElves(strings.NewReader(tt.input))
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v\n", err, tt.wantErr)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got = %v, want = %v\n", got, tt.want)
			}
		})
	}
}

func TestSumTop(t *testing.T) {
	tests := []struct {
		name  string
		elves []int
		n     int
		want  int
	}{
		{"sample 1", []int{4000, 6000, 10000, 11000, 24000}, 1, 24000},
		{"sample 2", []int{4000, 6000, 10000, 11000, 24000}, 3, 45000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumTop(tt.elves, tt.n)
			if got != tt.want {
				t.Errorf("got = %d, want = %d\n", got, tt.want)
			}
		})
	}
}
