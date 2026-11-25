package main

import (
	"maps"
	"strings"
	"testing"
)

const example = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

func exampleSizes() map[string]int {
	return dirSizes(strings.Split(example, "\n"))
}

func TestDirSizes(t *testing.T) {
	want := map[string]int{"/": 48381165, "/a": 94853, "/ae": 584, "/d": 24933642}
	if got := exampleSizes(); !maps.Equal(got, want) {
		t.Errorf("got = %v, want = %v", got, want)
	}
}

func TestParts(t *testing.T) {
	tests := []struct {
		name string
		part func(map[string]int) int
		want int
	}{
		{"part1", part1, 95437},
		{"part2", part2, 24933642},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.part(exampleSizes()); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
