package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"sort"
)

var adapters []int

var dynamicProgramming map[int]uint64

func findCombinations(idx int) uint64 {
	if dynamicProgramming[idx] != 0 {
		// fmt.Println("dynamicProgramming[", idx, "] = ", dynamicProgramming[idx])
		return dynamicProgramming[idx]
	}
	// fmt.Println("current index: ", idx)
	if idx == 0 { return 1 }

	// adapters[idx-1] always valid
	//ownFactor := uint64(1)
	recursiveFactor := findCombinations(idx-1)

	// adapters[idx-2] conditional
	if idx >= 2 && adapters[idx-2] >= adapters[idx] - 3 {
		//ownFactor++
		recursiveFactor += findCombinations(idx-2)
	}

	// adapters[idx-3] conditional
	if idx >= 3 && adapters[idx-3] >= adapters[idx] - 3 {
		//ownFactor++
		recursiveFactor += findCombinations(idx-3)
	}
	//res := ownFactor * recursiveFactor
	// fmt.Println("Completed index: ", idx, " (", adapters[idx], "); Result: ", recursiveFactor)
	dynamicProgramming[idx] = recursiveFactor
	return recursiveFactor
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	adapters = []int{0}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		this, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, this)
	}
	sort.Ints(adapters)
	dynamicProgramming = make(map[int]uint64)
	fmt.Println(findCombinations(len(adapters)-1))
}
