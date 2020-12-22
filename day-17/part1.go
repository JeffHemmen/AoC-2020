package main

import (
	"os"
	"fmt"
	"bufio"
)

type coord struct {
	x, y, z int
}

var min, max coord
// var world map[coord]rune

func expandWorld(cube coord) {
	if cube.x < min.x + 1 { min.x = cube.x - 1 }
	if cube.y < min.y + 1 { min.y = cube.y - 1 }
	if cube.z < min.z + 1 { min.z = cube.z - 1 }
	if cube.x > max.x - 1 { max.x = cube.x + 1 }
	if cube.y > max.y - 1 { max.y = cube.y + 1 }
	if cube.z > max.z - 1 { max.z = cube.z + 1 }
}

func isActive(world map[coord]rune, cube coord) bool {
	return world[cube] == '#'
}

func countNeighbours(world map[coord]rune, cube coord) int {
	res := 0
        for x := cube.x - 1; x <= cube.x + 1; x++ {
		for y := cube.y - 1; y <= cube.y + 1 ; y++ {
			for z := cube.z - 1; z <= cube.z + 1 ; z++ {
				neighbour := coord{x, y, z}
				if cube == neighbour { continue }
				if isActive(world, neighbour) { res++ }
			}
		}
	}
	return res
}


func countActive(world map[coord]rune) int{
	res := 0
        for x := min.x + 1; x <= max.x - 1; x++ {
                for y := min.y + 1; y <= max.y - 1; y++ {
                        for z := min.z + 1; z <= max.z - 1; z++ {
				cube := coord{x, y, z}
				if isActive(world, cube) {
					res++
				}
			}
		}
	}
	return res
}

func cycle(world map[coord]rune) map[coord]rune {
	newWorld := make(map[coord]rune)
	for k,v := range world {
		newWorld[k] = v
	}
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			for z := min.z; z <= max.z; z++ {
				cube := coord{x, y, z}
				numNeighbours := countNeighbours(world, cube)
				if isActive(world, cube) {
					if numNeighbours != 2 && numNeighbours != 3 {
						newWorld[cube] = '.'
					}
				} else { // inactive
					if numNeighbours == 3 {
						newWorld[cube] = '#'
						expandWorld(cube)
					}
				}
			}
		}
	}
	return newWorld
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	world := make(map[coord]rune)
	for x := 0; scanner.Scan(); x++ {
		for y, char := range scanner.Text() {
			world[coord{x, y, 0}] = char
			expandWorld(coord{x, y, 0})
		}
	}
	f.Close()

	for i := 0; i < 6; i++ {
		world = cycle(world)
	}

	fmt.Println(countActive(world))
}
