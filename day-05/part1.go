package main

import (
	"os"
	"fmt"
	"bufio"
	"math"
)

func decodeBinary(digits string, one rune) int {
	res := 0
	for i, p := range digits {
		if p == one {
			res += int(math.Pow(float64(2), float64(len(digits)-i-1)))
		}
	}
	return res
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	highscore := 0

	for scanner.Scan() {
		thisBSP := scanner.Text()
		thisSeat := decodeBinary(thisBSP[:7], 'B')
		thisSeat *= 8
		thisSeat += decodeBinary(thisBSP[7:], 'R')
		if thisSeat > highscore {
			highscore = thisSeat
		}
	}
	fmt.Println(highscore)
}
