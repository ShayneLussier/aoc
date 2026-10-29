package main

import (
	"os"
	"strings"
	"testing"
)

func TestTwoSum(t *testing.T) {
	actual, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		input  string
		target int
		want   int
	}{
		{"sample", "1721 979 366 299 675 1456", 2020, 514579},
		{"actual", string(actual), 2020, 326211},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := parseInput(strings.NewReader(tt.input))
			if err != nil {
				t.Fatal(err)
			}

			got := twoSum(input, tt.target)
			if got != tt.want {
				t.Errorf("got = %d, want = %d", got, tt.want)
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	actual, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		input  string
		target int
		want   int
	}{
		{"sample", "1721 979 366 299 675 1456", 2020, 241861950},
		{"actual", string(actual), 2020, 131347190},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := parseInput(strings.NewReader(tt.input))
			if err != nil {
				t.Fatal(err)
			}

			got := threeSum(input, tt.target)
			if got != tt.want {
				t.Fatalf("got = %d, want = %d", got, tt.want)
			}
		})
	}
}
