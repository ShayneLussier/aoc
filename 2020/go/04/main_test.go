package main

import (
	"strings"
	"testing"
)

func TestValidYear(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"2002", true},
		{"1920", true},
		{"2003", false},
		{"1919", false},
		{"02000", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := validYear(tt.in, 1920, 2002); got != tt.want {
			t.Errorf("validYear(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestValidHeight(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"60in", true},
		{"190cm", true},
		{"190in", false},
		{"190", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := validHeight(tt.in); got != tt.want {
			t.Errorf("validHeight(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestValidHairColor(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"#123abc", true},
		{"#123abz", false},
		{"123abc", false},
		{"#1zzzzz", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := validHairColor(tt.in); got != tt.want {
			t.Errorf("validHairColor(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestValidEyeColor(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"brn", true},
		{"wah", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := validEyeColor(tt.in); got != tt.want {
			t.Errorf("validEyeColor(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestValidPID(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"000000001", true},
		{"0123456789", false},
		{"1abc", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := validPID(tt.in); got != tt.want {
			t.Errorf("validPID(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestExamples(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		part1, part2 int
	}{
		{"part 1 example",
			`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`,
			2, 2},
		{"part 2 invalid",
			`eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`,
			4, 0},
		{"part 2 valid",
			`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`,
			4, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passports, err := parseInput(strings.NewReader(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			part1, part2 := parsePassports(passports)
			if part1 != tt.part1 || part2 != tt.part2 {
				t.Errorf("got part1=%d part2=%d, want part1=%d part2=%d", part1, part2, tt.part1, tt.part2)
			}
		})
	}
}
