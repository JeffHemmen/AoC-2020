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

	thisGroup := make(map[rune]bool)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			runningTotal += len(thisGroup)
			thisGroup = make(map[rune]bool)
		}
		for _, r := range line {
			thisGroup[r] = true
		}
	}
	runningTotal += len(thisGroup)
	fmt.Println(runningTotal)
}
