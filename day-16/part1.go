package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

type minmax struct {
	min, max int
}

var rules []minmax

func parseRule(line string) {
	// departure location: 32-842 or 854-967
	line = strings.Split(line, ": ")[1]
	// 32-842 or 854-967
	ranges := strings.Split(line, " or ")
	for _, rng := range ranges {
		mm := strings.Split(rng, "-")
		min, _ := strconv.Atoi(mm[0])
		max, _ := strconv.Atoi(mm[1])
		rules = append(rules, minmax{min, max})
	}
}

func isValid(num int) bool {
	for _, rule := range rules {
		if rule.min <= num && num <= rule.max {
			return true
		}
	}
	return false
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for _, line := scanner.Scan(), scanner.Text(); line != ""; _, line =  scanner.Scan(), scanner.Text() {
		parseRule(line)
	}

	scanner.Scan() // line that says "your ticket:"
	for _, line := scanner.Scan(), scanner.Text(); line != ""; _, line =  scanner.Scan(), scanner.Text() {
		// skip own ticket for now
	}

	total := 0

	scanner.Scan() // line that says "nearby tickets:"
	for _, line := scanner.Scan(), scanner.Text(); line != ""; _, line =  scanner.Scan(), scanner.Text() {
		strValues := strings.Split(line, ",")
		for _, strVal := range strValues {
			val, _ := strconv.Atoi(strVal)
			if !isValid(val) { total += val }
		}
	}
	fmt.Println(total)
}
