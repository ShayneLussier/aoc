// https://adventofcode.com/2021/day/4

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type board [5][5]int

type pos struct {
	board int
	row   int
	col   int
}

func parseInput(r io.Reader) (drawOrder []int, boards []board, where map[int][]pos, err error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	for _, num := range strings.Split(strings.TrimSpace(scanner.Text()), ",") {
		n, err := strconv.Atoi(num)
		if err != nil {
			return drawOrder, boards, where, err
		}
		drawOrder = append(drawOrder, n)
	}

	var b board
	where = map[int][]pos{}
	row := 0
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		for i, num := range fields {
			n, err := strconv.Atoi(num)
			if err != nil {
				return drawOrder, boards, where, err
			}
			b[row][i] = n
			where[n] = append(where[n], pos{board: len(boards), row: row, col: i})
		}
		row++
		if row == 5 {
			boards = append(boards, b)
			b = board{}
			row = 0
		}
	}
	return drawOrder, boards, where, scanner.Err()
}

func validBingo(b board, row, col int) bool {
	rowSum, colSum := 0, 0

	for i := range 5 {
		rowSum += b[row][i]
		colSum += b[i][col]
	}

	return rowSum == -5 || colSum == -5
}

func checkSum(b board) int {
	var sum int
	for _, r := range b {
		for _, n := range r {
			if n != -1 {
				sum += n
			}
		}
	}
	return sum
}

// playBingo computes the score each board the moment in first wins
func playBingo(drawOrder []int, boards []board, where map[int][]pos) (firstScore, lastScore int) {
	boardScore := map[int]int{}

	for _, num := range drawOrder {
		// replace all instance of the drawn number to -1 in boards
		for _, p := range where[num] {
			boards[p.board][p.row][p.col] = -1

			if validBingo(boards[p.board], p.row, p.col) {
				// skip board if it has already won
				if _, ok := boardScore[p.board]; ok {
					continue
				}
				// record a board's winning score
				boardScore[p.board] = num * checkSum(boards[p.board])
				// record the score if it is the first or last winning board
				switch len(boardScore) {
				case 1:
					firstScore = num * checkSum(boards[p.board])
				case len(boards):
					lastScore = num * checkSum(boards[p.board])
				}
			}
		}
	}
	return firstScore, lastScore
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	drawOrder, boards, where, err := parseInput(f)
	if err != nil {
		log.Fatal(err)
	}

	firstScore, lastScore := playBingo(drawOrder, boards, where)
	fmt.Printf("part 1 = %d\n", firstScore) // 35711
	fmt.Printf("part 2 = %d\n", lastScore)  // 5586
}
