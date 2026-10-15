package main

import (
	"slices"
	"strings"
	"testing"
)

func TestParseMoves(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []move
		wantErr bool
	}{
		{"part 1 sample",
			"forward 5\ndown 5\nforward 8",
			[]move{{dir: "forward", n: 5}, {dir: "down", n: 5}, {dir: "forward", n: 8}},
			false},
		{"not a number", "forward hello", []move{}, true},
		{"blank line", "up 1\n\nup 2", []move{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMoves(strings.NewReader(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcDepth(t *testing.T) {
	tests := []struct {
		name  string
		input []move
		want  int
	}{
		{"part 1 sample",
			[]move{
				{dir: "forward", n: 5},
				{dir: "down", n: 5},
				{dir: "forward", n: 8},
				{dir: "up", n: 3},
				{dir: "down", n: 8},
				{dir: "forward", n: 2},
			},
			150},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcDepth(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcAim(t *testing.T) {
	tests := []struct {
		name  string
		input []move
		want  int
	}{
		{"part 2 sample",
			[]move{
				{dir: "forward", n: 5},
				{dir: "down", n: 5},
				{dir: "forward", n: 8},
				{dir: "up", n: 3},
				{dir: "down", n: 8},
				{dir: "forward", n: 2},
			},
			900},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcAim(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
