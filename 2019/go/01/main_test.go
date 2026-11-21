package main

import "testing"

func TestFuelRequirement(t *testing.T) {
	tests := []struct {
		mass, want int
	}{
		{14, 2},
		{12, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, tt := range tests {
		if got := fuelRequirement(tt.mass); got != tt.want {
			t.Errorf("mass %d: got = %d, wanted = %d", tt.mass, got, tt.want)
		}
	}
}

func TestFuelSum(t *testing.T) {
	tests := []struct {
		input             []int
		simple, recursive int
	}{
		{[]int{12, 14, 1969, 100756}, 34241, 51316},
	}
	for _, tt := range tests {
		simple, recursive := fuelSum(tt.input)
		if simple != tt.simple || recursive != tt.recursive {
			t.Errorf("%v: got = %d, %d, wanted %d, %d\n", tt.input, simple, recursive, tt.simple, tt.recursive)
		}
	}
}
