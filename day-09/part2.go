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
	var invalidNum int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		XMAS = append(XMAS, i)
	}

	for i := preamble; i < len(XMAS); i++ {
		if !isCompliant(i) {
			invalidNum = XMAS[i]
			break
		}
	}

	for i := 0; i < len(XMAS); i++ {
		counter := 0
		smallest, largest := XMAS[i], XMAS[i]
		for j := i; counter < invalidNum; j++ {
			counter += XMAS[j]
			if XMAS[j] < smallest { smallest = XMAS[j] }
			if XMAS[j] > largest  { largest  = XMAS[j] }
		}
		if counter == invalidNum && smallest != largest {
			fmt.Println(smallest + largest)
		}
	}

}
