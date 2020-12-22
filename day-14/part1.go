package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"strings"
	"regexp"
)

var andMask, orMask uint64
var memory map[uint64]uint64

func parseMask(line string) {
	mask := strings.Fields(line)[2]
	// andMask = ^uint64(0)	// 64 1's
	andMask = (1 << 36) - 1 // 36 1's
	orMask  = 0		// 64 0's
	var thisBit uint64
	for i, b := range mask {
		if b == '1' {
			thisBit = 1 << (35 - i)
			orMask |= thisBit
		} else if b == '0' {
			thisBit = 1 << (35 - i)
			andMask ^= thisBit
		}
	}
}

func parseMem(line string) {
	rgx, _ := regexp.Compile("^mem\\[([0-9]+)\\] = ([0-9]+)$")
	match := rgx.FindStringSubmatch(line)
	if match == nil { panic("no regex match") }
	strAddr, strVal := match[1], match[2]
	addr, _ := strconv.ParseUint(strAddr, 10, 64)
	val,  _ :=strconv.ParseUint(strVal, 10, 64)
	val &= andMask
	val |= orMask
	memory[addr] = val
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	memory = make(map[uint64]uint64)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			parseMask(line)
		} else {
			parseMem(line)
		}
	}

	var result uint64 = 0
	for _, v := range memory {
		result += v
	}
	fmt.Println(result)
}
