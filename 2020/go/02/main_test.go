package main

import (
	"os"
	"slices"
	"strings"
	"testing"
)

const sampleInput = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestParseInput(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		wantPolicies  []policy
		wantPasswords []string
		wantErr       bool
	}{
		{
			name:  "sample",
			input: sampleInput,
			wantPolicies: []policy{
				{min: 1, max: 3, val: 'a'},
				{min: 1, max: 3, val: 'b'},
				{min: 2, max: 9, val: 'c'},
			},
			wantPasswords: []string{"abcde", "cdefg", "ccccccccc"},
		},
		{
			name:    "bad input",
			input:   "nothing here!",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPolicies, gotPasswords, gotErr := parseInput(strings.NewReader(tt.input))
			if (gotErr != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr = %v\n", gotErr, tt.wantErr)
			}

			// if we get the expected error
			if gotErr != nil {
				return
			}

			if !slices.Equal(gotPolicies, tt.wantPolicies) {
				t.Errorf("polices = %v, want = %v\n", gotPolicies, tt.wantPolicies)
			}

			if !slices.Equal(gotPasswords, tt.wantPasswords) {
				t.Errorf("passwords = %v, want = %v\n", gotPasswords, tt.wantPasswords)
			}
		})
	}
}

func TestValidPasswords(t *testing.T) {
	policies, passwords, err := parseInput(strings.NewReader(sampleInput))
	if err != nil {
		t.Fatal(err)
	}

	got := validPasswords(policies, passwords)
	want := 2

	if got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

func TestActualValidPasswords(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	policies, passwords, err := parseInput(f)
	if err != nil {
		t.Fatal(err)
	}

	got := validPasswords(policies, passwords)
	want := 550

	if got != want {
		t.Errorf("got = %d, want = %d\n", got, want)
	}
}

func TestValidPasswords2(t *testing.T) {
	policies, passwords, err := parseInput(strings.NewReader(sampleInput))
	if err != nil {
		t.Fatal(err)
	}

	got := validPasswords2(policies, passwords)
	want := 1

	if got != want {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

func TestActualValidPasswords2(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	policies, passwords, err := parseInput(f)
	if err != nil {
		t.Fatal(err)
	}

	got := validPasswords2(policies, passwords)
	want := 634

	if got != want {
		t.Errorf("got = %d, want = %d\n", got, want)
	}
}
