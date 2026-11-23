package main

import "testing"

func TestIntersection(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]byte
		want   byte
		wantOk bool
	}{
		{"sample1", [][]byte{[]byte("vJrwpWtwJgWr"), []byte("hcsFMMfFFhFp")}, 'p', true},
		{"sample2", [][]byte{[]byte("jqHRNqRjqzjGDLGL"), []byte("rsFMfFZSrLrFZsSL")}, 'L', true},
		{"sample3", [][]byte{[]byte("PmmdzqPrV"), []byte("vPwwTWBwg")}, 'P', true},
		{"sample4", [][]byte{[]byte("vJrwpWtwJgWrhcsFMMfFFhFp"), []byte("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"), []byte("PmmdzqPrVvPwwTWBwg")}, 'r', true},
		{"sample5", [][]byte{[]byte("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"), []byte("ttgJtRGJQctTZtZT"), []byte("CrZsJsPPZsGzwwsLwLmpwMDw")}, 'Z', true},
		{"empty input", [][]byte{[]byte("")}, 0, false},
		{"no input", nil, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := intersection(tt.input...)
			if ok != tt.wantOk {
				t.Fatalf("ok = %v, wantOk = %v\n", ok, tt.wantOk)
			}
			if got != tt.want {
				t.Errorf("got = %v, want = %v\n", got, tt.want)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int
		wantErr bool
	}{
		{"test1",
			[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
			157,
			false},
		{"lowest lowercase", []string{"aa"}, 1, false},
		{"highest lowercase", []string{"zz"}, 26, false},
		{"lowest uppercase", []string{"AA"}, 27, false},
		{"highest uppercase", []string{"ZZ"}, 52, false},
		{"no common item", []string{"ab"}, 0, true},
		{"empty line", []string{""}, 0, true},
		{"no input", nil, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got = %d, want = %d", got, tt.want)
			}
		})
	}
}

func TestParseBadges(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int
		wantErr bool
	}{
		{"sample", []string{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		}, 70, false},
		{"lowest lowercase", []string{"a", "a", "a"}, 1, false},
		{"highest lowercase", []string{"z", "z", "z"}, 26, false},
		{"lowest uppercase", []string{"A", "A", "A"}, 27, false},
		{"highest uppercase", []string{"Z", "Z", "Z"}, 52, false},
		{"no common item", []string{"a", "b", "c"}, 0, true},
		{"common to two of three", []string{"a", "a", "b"}, 0, true},
		{"empty lines", []string{"", "", ""}, 0, true},
		{"incomplete group ignored", []string{"a", "a", "a", "b", "b"}, 1, false},
		{"no input", nil, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBadges(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got = %d, want = %d", got, tt.want)
			}
		})
	}
}
