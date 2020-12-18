package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"sort"
	"math/big"
)

var adapters []int

var dynamicProgramming map[int]big.Int

func findCombinations(idx int) big.Int {
	cache := dynamicProgramming[idx]
	if cache.Cmp(new(big.Int)) != 0 {
		fmt.Println("dynamicProgramming[", idx, "] = ", dynamicProgramming[idx])
		return dynamicProgramming[idx]
	}
	fmt.Println("current index: ", idx)
	if idx == 0 { return *new(big.Int).SetInt64(int64(1)) }

	// adapters[idx-1] always valid
	ownFactor := 1
	recursiveFactor := findCombinations(idx-1)

	// adapters[idx-2] conditional
	if idx >= 2 && adapters[idx-2] >= adapters[idx] - 3 {
		ownFactor++
		recursiveFactor.Mul(&recursiveFactor, findCombinations(idx-2))
	}

	// adapters[idx-3] conditional
	if idx >= 3 && adapters[idx-3] >= adapters[idx] - 3 {
		ownFactor++
		recursiveFactor *= findCombinations(idx-3)
	}
	res := big.Mul(ownFactor, recursiveFactor)
	fmt.Println("Completed index: ", idx, " (", adapters[idx], "); Result: ", res)
	dynamicProgramming[idx] = res
	return res
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
	dynamicProgramming = make(map[int]big.Int)
	fmt.Println(findCombinations(20))
}
