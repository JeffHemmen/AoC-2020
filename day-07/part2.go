package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

const shinygold = "shiny gold"

type Bag struct {
	colour string
	requiredBags  []string
	requiredCount []int
}

var bagIndex map[string]Bag

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

func countBags(colour string) (count int) {
	count = 1
	bag := bagIndex[colour]
	for i, _ := range bag.requiredBags {
		count += bag.requiredCount[i] * countBags(bag.requiredBags[i])
	}
	return
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	bagIndex = make(map[string]Bag)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		thisBag := parseRule(scanner.Text())
		bagIndex[thisBag.colour] = thisBag
	}

	fmt.Println(countBags(shinygold) - 1)
}
