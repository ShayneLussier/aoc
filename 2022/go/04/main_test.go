package main

import (
	"io"
	"slices"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input io.Reader
		want  []pair
	}{
		{"example1", strings.NewReader("8-41,8-79"), []pair{{{lo: 8, hi: 41}, {lo: 8, hi: 79}}}},
		{"example2", strings.NewReader("1-71,2-71"), []pair{{{lo: 1, hi: 71}, {lo: 2, hi: 71}}}},
		{"example3", strings.NewReader("11-74,74-75"), []pair{{{lo: 11, hi: 74}, {lo: 74, hi: 75}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.input)
			if err != nil {
				t.Fatalf("err = %v\n", err)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got = %v, want = %v\n", got, tt.want)
			}
		})
	}
}

func TestParseSections(t *testing.T) {
	tests := []struct {
		name  string
		input []pair
		want  int
	}{
		{"example1", []pair{{{lo: 8, hi: 41}, {lo: 8, hi: 79}}}, 1},
		{"example2", []pair{{{lo: 1, hi: 71}, {lo: 2, hi: 71}}}, 1},
		{"example3", []pair{{{lo: 11, hi: 74}, {lo: 74, hi: 75}}}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseSections(tt.input)
			if got != tt.want {
				t.Errorf("got = %d, want = %d\n", got, tt.want)
			}
		})
	}
}

func TestParseDections2(t *testing.T) {
	tests := []struct {
		name  string
		input []pair
		want  int
	}{
		{"example1", []pair{{{lo: 8, hi: 41}, {lo: 8, hi: 79}}}, 1},
		{"example2", []pair{{{lo: 1, hi: 71}, {lo: 2, hi: 71}}}, 1},
		{"example3", []pair{{{lo: 11, hi: 74}, {lo: 74, hi: 75}}}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseSections2(tt.input)
			if got != tt.want {
				t.Errorf("got = %d, want = %d\n", got, tt.want)
			}
		})
	}
}
