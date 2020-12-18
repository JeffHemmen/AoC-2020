package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

const shinygold = "shiny golden"

type Bag struct {
	colour string
	requiredBags  []string
	requiredCount []int
}

func parseRule(line string) Bag {
	var ruleBag string
	var requiredBags []string = make([]string, 0)
	var requiredCount []int = make([]int, 0)

	fields := strings.Fields(line)

	ruleBag = fields[0] + " " + fields[1]

	for i := 4; i < len(fields); i += 4 {
		c, _ := strconv.Atoi(fields[i])
		requiredCount = append(requiredCount, c)
		requiredBags  = append(requiredBags, fields[i+1] + " " + fields[i+2])
	}

	return Bag{ruleBag, requiredBags, requiredCount}
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	bagIndex := make(map[string]Bag)
	entrypointFinder := make(map[string]bool)
	entrypoints := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		thisBag := parseRule(scanner.Text())
		bagIndex[thisBag.colour] = thisBag
		entrypointFinder[thisBag.colour] = true
		// fmt.Println(thisBag)
	}

	// find DAG entrypoint
	for _, bag := range bagIndex {
		for _, eliminated := range bag.requiredBags {
			entrypointFinder[eliminated] = false
			fmt.Println("eliminating " + eliminated)
		}
	}
	for entrypoint, candidate := range entrypointFinder {
		if candidate {
			entrypoints = append(entrypoints, entrypoint)
		}
	}
	
	// iterate DAG from every entrypoint
	// - if entrypoint is shinygold, skip
	// - if entrypoint does not contain shinygold, do nothing

}
