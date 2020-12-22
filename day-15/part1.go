package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

const limit = 202000

func naive(nums []int) {
	var i int
	for i = len(nums); i < limit; i++ {
		last := nums[i-1]
		var j int
		for j = i-2; j >= 0; j-- {
			if nums[j] == last {
				// prior occurrence at index j
				nums = append(nums, (i - 1) - j)
				break
			}
		}
		if j < 0 {
			// no prior occurrence found
			nums = append(nums, 0)
		}
	}
	fmt.Println(nums[i-1])
}

func smarter(nums []int) {

	occ := make(map[int]int)

	for i, n := range nums {
		occ[n] = i+1
	}

	next := nums[len(nums) - 1]
	occ[next] = 0

	for i := len(nums); i < limit; i++ {
		prevIdx := occ[next]
		occ[next] = i
		if prevIdx == 0 {
			next = 0
		} else {
			next = i - prevIdx
		}
	}
	fmt.Println(next)
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	f.Close()
	line := scanner.Text()
	startingNumbersStr := strings.Split(line, ",")
	startingNumbers := make([]int, 0)
	for _, n := range startingNumbersStr {
		thisNum, _ := strconv.Atoi(n)
		startingNumbers = append(startingNumbers, thisNum)
	}

	// naive(startingNumbers)
	smarter(startingNumbers)

	/* With an increased limit of 202000,
	   these are my benchmark findings
	   on a MacBook Pro (13-inch, 2018):

	     ./naive    3.86s user 0.11s system 111% cpu 3.557 total
	     ./smarter  0.01s user 0.00s system 91%  cpu 0.015 total
	*/
}
