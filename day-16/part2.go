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

var class map[string][]minmax


func parseRule(line string) {
	// line = departure location: 32-842 or 854-967
	parts := strings.Split(line, ": ")
	className, catRanges := parts[0], parts[1]
	// catRanges 32-842 or 854-967
	ranges := strings.Split(catRanges, " or ")
	theseRules := make([]minmax, 0)
	for _, rng := range ranges {
		mm := strings.Split(rng, "-")
		min, _ := strconv.Atoi(mm[0])
		max, _ := strconv.Atoi(mm[1])
		theseRules = append(theseRules, minmax{min, max})
	}
	class[className] = theseRules
}

func isValidForClass(className string, num int) bool {
	for _, rule := range class[className] {
		if rule.min <= num && num <= rule.max {
			return true
		}
	}
	return false
}

func isValidForAny(num int) bool {
	for className, _ := range class {
		if isValidForClass(className, num) { return true }
	}
	return false
}

func del(slice []int, idx int) []int {
	slice[idx] = slice[0]
	return slice[1:]
}

func findAndRemove(slice []int, elem int) ([]int, bool) {
	for i, e := range slice {
		if e != elem { continue }
		return del(slice, i), true
	}
	return slice, false
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	class = make(map[string][]minmax)

	for _, line := scanner.Scan(), scanner.Text(); line != ""; _, line =  scanner.Scan(), scanner.Text() {
		parseRule(line)
	}

	myTicket := make([]int, 0)
	scanner.Scan() // line that says "your ticket:"
	for _, line := scanner.Scan(), scanner.Text(); line != ""; _, line =  scanner.Scan(), scanner.Text() {
		strValues := strings.Split(line, ",")
		for _, strVal := range strValues {
			val, _ := strconv.Atoi(strVal)
			myTicket = append(myTicket, val)
		}
	}

	nearbyTickets := make([][]int, 0)

	scanner.Scan() // line that says "nearby tickets:"
	for _, line := scanner.Scan(), scanner.Text(); line != ""; _, line =  scanner.Scan(), scanner.Text() {
		validTicket := true
		thisTicket := make([]int, 0)
		strValues := strings.Split(line, ",")
		for _, strVal := range strValues {
			val, _ := strconv.Atoi(strVal)
			if !isValidForAny(val) {
				validTicket = false
				break
			}
			thisTicket = append(thisTicket, val)
		}
		if !validTicket { continue }
		nearbyTickets = append(nearbyTickets, thisTicket)
	}


	transposedTickets := make([][]int, len(nearbyTickets[0]))
	for i, _ := range transposedTickets {
		transposedTickets[i] = make([]int, 0)
	}

	for _, nearbyTicket := range nearbyTickets {
		for i, val := range nearbyTicket {
			transposedTickets[i] = append(transposedTickets[i], val)
		}
	}

	possComb := make(map[string][]int)
	for className, _ := range class {
		possComb[className] = make([]int, 0)
	}

	for i, col := range transposedTickets {
		for className, _ := range class {
			valid := true
			for _, val := range col {
				if !isValidForClass(className, val) {
					valid = false
					break
				}
			}
			if valid { // all values in column okay for this class
				possComb[className] = append(possComb[className], i)
			}
		}
	}

	for unchanged, allFound := false, false; !unchanged && !allFound; {
		unchanged, allFound = true, true
		for className, poss := range possComb {
			if len(poss) > 1 {
				allFound = false
				continue
			}
			// len(poss) == 1
			for otherClass, otherPoss := range possComb {
				if otherClass == className { continue }
				var found bool
				possComb[otherClass], found = findAndRemove(otherPoss, possComb[className][0])
				if found { unchanged = false }
			}
		}
	}

	result := 1
	for className, poss := range possComb {
		if !strings.HasPrefix(className, "departure") { continue }
		result *= myTicket[poss[0]]
	}
	fmt.Println(result)
}
