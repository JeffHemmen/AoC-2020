package main

import (

	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func op(line []string, i int) string {
	a, op, b := line[i], line[i+1], line[i+2]
	inta, _ := strconv.Atoi(a)
	intb, _ := strconv.Atoi(b)
	var intres int
	if op == "*" {
		intres = inta * intb
	} else {
		intres = inta + intb
	}
	res := fmt.Sprintf("%d", intres)
	return res
}

func deleteIdx(line *[]string, idx int) {
        *line = append((*line)[:idx], (*line)[idx+1:]...)
}

func resolveBracket(line *[]string, start, end int) {
	numLB, numRB := strings.Count((*line)[start], "("), strings.Count((*line)[end], ")")
	(*line)[start] = strings.Trim((*line)[start], "(")
	(*line)[end]   = strings.Trim((*line)[end], ")")
	allAdditionsDone := false
	for i := start; i <= end; i++ {
		if (*line)[i] == "+" {
			(*line)[i-1] = op(*line, i-1)
			deleteIdx(line, i+1)
			deleteIdx(line, i)
			end -= 2
			i = start - 1 // will be incremented to `start` on next iteration
			continue
		}
		if i == end && !allAdditionsDone {
			allAdditionsDone = true
			i = start
		}
		if allAdditionsDone && (*line)[i] == "*" {
			(*line)[i-1] = op(*line, i-1)
			deleteIdx(line, i+1)
			deleteIdx(line, i)
			end -= 2
			i = start - 1 // will be incremented to `start` on next iteration
		}
	}
	// add brackets back
	for i := 0; i < numLB-1; i++ {
		(*line)[end] = "(" + (*line)[end]
	}
	for i := 0; i < numRB-1; i++ {
		(*line)[end] = (*line)[end] + ")"
	}
}

func parseLine(line []string) int {
	// get rid of all brackets first
	var start int
	for i := 0; i < len(line); i++ {
		if strings.HasPrefix(line[i], "(") {
			start = i
		} else if strings.HasSuffix(line[i], ")") {
			resolveBracket(&line, start, i)
			i = -1 // will be incremented to `0` on next iteration
		}
	}
	resolveBracket(&line, 0, len(line)-1)
	res, _ := strconv.Atoi(line[0])
	return res
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	total, res := 0, 0
        for scanner.Scan() {
                line := scanner.Text()
                lineFields := strings.Fields(line)
                res = parseLine(lineFields)
                total += res
        }

        fmt.Println(total)
}
