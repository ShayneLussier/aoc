package main

import "testing"

func TestScore1(t *testing.T) {
	tests := []struct {
		j, k byte
		want int
	}{
		{'A', 'X', 4},
		{'A', 'Y', 8},
		{'A', 'Z', 3},
		{'B', 'X', 1},
		{'B', 'Y', 5},
		{'B', 'Z', 9},
		{'C', 'X', 7},
		{'C', 'Y', 2},
		{'C', 'Z', 6},
		{'D', 'D', 0},
	}

	for _, tt := range tests {
		got := scorePart1(tt.j, tt.k)
		if got != tt.want {
			t.Errorf("got = %d, want = %d\n", got, tt.want)
		}
	}
}

func TestScore2(t *testing.T) {
	tests := []struct {
		j, k byte
		want int
	}{
		{'A', 'X', 3},
		{'A', 'Y', 4},
		{'A', 'Z', 8},
		{'B', 'X', 1},
		{'B', 'Y', 5},
		{'B', 'Z', 9},
		{'C', 'X', 2},
		{'C', 'Y', 6},
		{'C', 'Z', 7},
		{'D', 'D', 0},
	}

	for _, tt := range tests {
		got := scorePart2(tt.j, tt.k)
		if got != tt.want {
			t.Errorf("got = %d, want = %d\n", got, tt.want)
		}
	}
}
