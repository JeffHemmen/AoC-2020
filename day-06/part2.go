package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	runningTotal := 0

	thisGroup := make(map[rune]int)
	scanner := bufio.NewScanner(f)
	groupSize := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for _, c := range thisGroup {
				if c == groupSize {
					runningTotal++
				}
			}
			thisGroup = make(map[rune]int)
			groupSize = 0
			continue
		}
		groupSize++
		for _, r := range line {
			thisGroup[r] += 1
		}

	}
	for _, c := range thisGroup {
		if c == groupSize {
			runningTotal++
		}
	}
	fmt.Println(runningTotal)
}
