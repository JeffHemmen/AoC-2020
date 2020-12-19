package main

import (
	"os"
	"bufio"
	"fmt"
)

var seats [][]rune

func exists(seats [][]rune, row, col int) bool {
	if row < 0 || col < 0 || row >= len(seats) || col >= len(seats[0]) {
		return false
	}
	return true
}


func isFloor(seats [][]rune, row, col int) bool {
	if seats[row][col] == '.' { return true }
	return false
}

func isOccupied(seats [][]rune, row, col int) int {
	if !exists(seats, row, col) {
		return 0
	}
	if seats[row][col] == '#' { return 1 }
	return 0
}

func countVisible(seats [][]rune, from_row, from_col int) (counter int) {
	for rowDir := -1; rowDir <= 1; rowDir++ {
		for colDir := -1; colDir <= 1; colDir++ {
			if rowDir == 0 && colDir == 0 { continue }
			for dist := 1; true; dist++ {
				row, col := dist*rowDir+from_row, dist*colDir+from_col
				if !exists(seats, row, col) {
					break
				}
				if isFloor(seats, row, col) {
					continue
				}
				counter += isOccupied(seats, row, col)
				break
			}
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
			if seats[row][col] == 'L' && countVisible(seats, row, col) == 0 {
				newRow = append(newRow, '#')
				changed = true
			} else if seats[row][col] == '#' && countVisible(seats, row, col) >= 5 {
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
