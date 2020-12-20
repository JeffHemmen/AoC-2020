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

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	east, south := 0, 0
	orientation := 0 //east=0, south=1, west=2, north=3

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		action, value := parseInstruction(scanner.Text())
		switch action {
			case 'N':
				south -= value
			case 'S':
				south += value
			case 'E':
				east += value
			case 'W':
				east -= value
			case 'L':
				orientation = (orientation + (360-value)/90) % 4
			case 'R':
				orientation = (orientation + value/90) % 4
			case 'F':
				switch orientation {
				case 0: // east
					east += value
				case 1: // south
					south += value
				case 2: // west
					east -= value
				case 3: // north
					south -= value
				}
		}
		fmt.Printf("%c%d =>  E%dS%d\n", action, value, east, south)
	}


	if east < 0 { east = - east }
	if south < 0 { south = - south }
	fmt.Println(east + south)
}
