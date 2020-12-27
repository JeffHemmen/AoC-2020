package main

import (

	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func parseOperator(line []string, idx *int) rune {
	*idx++
	return rune(line[*idx-1][0])
}

func parseOperand(line []string, idx *int) (uint64, int) {
	// fmt.Printf("    parseOperand(line, idx=%d)\n", *idx)
	thisField := line[*idx]
	res, err := strconv.ParseUint(thisField, 10, 64)
	if err == nil {
		*idx++
		// fmt.Printf("    returns %d, %d with idx=%d\n", res, 0, *idx)
		return res, 0
	} else if thisField[0] == '(' {
		line[*idx] = thisField[1:] // remove leading '('
		res, cb := parseLine(line, idx) // *idx will be incremented in here
		// fmt.Printf("    returns %d, %d with idx=%d\n", res, cb, *idx)
		return res, cb
	} else if strings.HasSuffix(thisField, ")") {
		// remove tailing ')' s
		cb := strings.Count(thisField, ")")
		thisField = strings.Trim(thisField, ")")
		res, _ := strconv.ParseUint(thisField, 10, 64)
		*idx++
		// fmt.Printf("    returns %d, %d with idx=%d\n", res, cb, *idx)
		return res, cb
		// what to do about idx? and the brackets?
	} else {
		panic("Could not parse")
	}
}

func parseLine(line []string, idx *int) (uint64, int) {
	if idx == nil {
		idx_val := 0
		idx = &idx_val
	}
	// fmt.Printf("parseLine(line, idx=%d)\n", *idx)
	a, cb := parseOperand(line, idx)
	if cb != 0 { panic("cb expected to be 0 here") }
	for *idx < len(line) && cb == 0 {
		var b uint64
		op := parseOperator(line, idx)
		b, cb  = parseOperand(line, idx)
		if op == '+' {
			a += b
		} else {
			a *= b
		}
	}
	if cb > 0 { cb-- }
	// fmt.Printf("  returns: %d, %d\n", a, cb)
	return a, cb
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)


	/*
	line := "9 + 5 * 6 * 7 * (2 * 8 * (2 * 7) * 5) * 7"
	fmt.Println(line, "\n")
	lineFields := strings.Fields(line)
	e, _ := parseLine(lineFields, nil)

	fmt.Println(e)
	*/


	total, res := uint64(0), uint64(0)
        for scanner.Scan() {
                line := scanner.Text()
                lineFields := strings.Fields(line)
                res, _ = parseLine(lineFields, nil)
                total += res
        }

        fmt.Println(total)



}
