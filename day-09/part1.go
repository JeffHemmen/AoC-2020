package main

import (
	"os"
	"fmt"
	"strconv"
	"bufio"
)

const preamble = 25

var XMAS []int

func isCompliant(idx int) bool {
	for i := idx-preamble; i<idx; i++ {
		for j := i+1; j <idx; j++ {
			if XMAS[i] + XMAS[j] == XMAS[idx] {
				return true
			}
		}
	}
	return false
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	XMAS = make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		XMAS = append(XMAS, i)
	}

	for i := preamble; i < len(XMAS); i++ {
		if ! isCompliant(i) {
			fmt.Println(XMAS[i])
			break
		}
	}
}
