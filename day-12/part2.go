package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func parseInstruction(raw string) (action rune, value int) {
	action = rune(raw[0])
	value, _ = strconv.Atoi(raw[1:])
	return
}

func matrixMul(mat1, mat2 [][]int) [][]int {
	m, n, p := len(mat1), len(mat2), len(mat2[0])
	result := make([][]int, m)
	for i:= 0; i<m; i++ {
		// fmt.Printf("  i=%d\n", i)
		result[i] = make([]int, p)
		for j := 0; j < p; j++ {
			// fmt.Printf("    j=%d\n", j)
			result[i][j] = 0
			for c := 0; c < n; c++ {
				// fmt.Printf("      c=%d\n", c)
				result[i][j] += mat1[i][c] * mat2[c][j]
			}
		}
	}
	return result
}

func sin(th int) int {
	// simplified sin, to work with my manual implementation
	th = (th + 360) % 360
	switch th {
		case 0  : return 0
		case 90 : return 1
		case 180: return 0
		case 270: return -1
		default : return 2 // impossible value
	}
}

func cos(th int) int {
	return sin(th+90)
}

func rotate(x, y, th int) (rx, ry int) {
	A    := [][]int{{x},{y}}
	rotM := [][]int{{cos(th), -sin(th)}, {sin(th), cos(th)}}
	AR   := matrixMul(rotM, A)
	rx,ry = AR[0][0], AR[1][0]
	return
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	shipEast, shipSouth := 0, 0
	wayEast, waySouth := 10, -1

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		action, value := parseInstruction(scanner.Text())
		switch action {
			case 'N':
				waySouth -= value
			case 'S':
				waySouth += value
			case 'E':
				wayEast += value
			case 'W':
				wayEast -= value
			case 'L':
				// rotate -value for L (normally positive) because out SOUTH-axis points down
				wayEast, waySouth = rotate(wayEast, waySouth, -value)
			case 'R':
				// converse of the above
				wayEast, waySouth = rotate(wayEast, waySouth, value)
			case 'F':
				shipEast += value * (wayEast)
				shipSouth += value * (waySouth)
		}
		// fmt.Printf("%c%d =>  E%d S%d (E%d S%d)\n", action, value, shipEast, shipSouth, wayEast, waySouth)
	}


	if shipEast < 0 { shipEast = - shipEast }
	if shipSouth < 0 { shipSouth = - shipSouth }
	fmt.Println(shipEast + shipSouth)
}
