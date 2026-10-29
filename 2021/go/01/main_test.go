package main

import (
	"slices"
	"strings"
	"testing"
)

func TestCountIncreases(t *testing.T) {
	sample := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	tests := []struct {
		name    string
		input   []int
		window  int
		want    int
		wantErr bool
	}{
		{"part1 sample", sample, 1, 7, false},
		{"part2 sample", sample, 3, 5, false},
		{"empty", nil, 1, 0, true},
		{"shorter than window", []int{1, 2}, 3, 0, true},
		{"zero window", sample, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countIncreases(tt.input, tt.window)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestReadInts(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{"valid", "199\n200\n208", []int{199, 200, 208}, false},
		{"empty", "", nil, false},
		{"not a number", "12\nabc", nil, true},
		{"blank line", "12\n\n13", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readInts(strings.NewReader(tt.input))
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tt.wantErr)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
