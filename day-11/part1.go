package main

import (
	"os"
	"bufio"
	"fmt"
)

var seats [][]rune

func isOccupied(seats [][]rune, row, col int) int {
	if row < 0 || col < 0 || row >= len(seats) || col >= len(seats[0]) {
		return 0
	}
	if seats[row][col] == '#' { return 1 }
	return 0
}

func countAdjacent(seats [][]rune, row, col int) (counter int) {
	for r := row-1; r <= row+1; r++ {
		for c := col-1; c <= col+1; c++ {
			if r == row && c == col { continue }
			counter += isOccupied(seats, r, c)
		}
	}
	return
}

func step(seats [][]rune) ([][]rune, bool) {
	newSeats := make([][]rune, 0)
	changed := false
	for row, _ := range seats {
		newRow := make([]rune, 0)
		for col, _ := range seats[row] {
			if seats[row][col] == 'L' && countAdjacent(seats, row, col) == 0 {
				newRow = append(newRow, '#')
				changed = true
			} else if seats[row][col] == '#' && countAdjacent(seats, row, col) >= 4 {
				newRow = append(newRow, 'L')
				changed = true
			} else {
				newRow = append(newRow, seats[row][col])
			}
		}
		newSeats = append(newSeats, newRow)
	}
	return newSeats, changed
}

func countOccupied(seats [][]rune) (counter int) {
	for _, row := range seats {
		for _, s := range row {
			if s == '#' {
				counter++
			}
		}
	}
	return
}

func printSeats(seats [][]rune) {
	for _, row := range seats {
		for _, s := range row {
			fmt.Printf("%c", s)
		}
		fmt.Println()
	}
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	seats = make([][]rune, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := make([]rune, 0)
		for _, s := range scanner.Text() {
			row = append(row, s)
		}
		seats = append(seats, row)
	}

	for changed := true; changed; seats, changed = step(seats) {
		// printSeats(seats)
	}

	fmt.Println(countOccupied(seats))

}
