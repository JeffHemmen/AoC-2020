package main

import (
	"os"
	"fmt"
	"bufio"
)

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

	var height, width int = len(grid), len(grid[1])

	var trees int
	for x, y := 0, 0; y < height; x, y = (x+3)%width, y+1 {
		if grid[y][x] { trees++ }
	}

	fmt.Println(trees)




}
