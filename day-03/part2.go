package main

import (
	"os"
	"fmt"
	"bufio"
)

var width, height int

func evalSlope(grid [][]bool, sx, sy int) (trees int) {
	for x, y := 0, 0; y < height; x, y = (x+sx)%width, y+sy {
		if grid[y][x] { trees++ }
	}
	return
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var grid [][]bool

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var row []bool
		for _, sq := range scanner.Text() {
			if sq == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		grid = append(grid, row)
	}

	height, width = len(grid), len(grid[1])

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var result int = 1
	for _, slope := range slopes {
		result *= evalSlope(grid, slope[0], slope[1])
	}

	fmt.Println(result)




}
