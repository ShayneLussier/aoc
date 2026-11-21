package main

import (
	"slices"
	"testing"
)

func TestParseCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []int
		want  []int
	}{
		{"example 1", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{"example 2", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{"example 3", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"example 4", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseCodes(tt.codes)
			if !slices.Equal(tt.codes, tt.want) {
				t.Errorf("got = %v, want = %v", tt.codes, tt.want)
			}
		})
	}
}
