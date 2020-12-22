package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"strings"
	"regexp"
)

var mask string
var memory map[uint64]uint64

func writeMem(addr, val uint64, mask string, offset int) {
	var thisBit uint64
	for i := offset; i < 36; i++ {
		if mask[i] == '0' {
			continue
		}
		thisBit = 1 << (35 - i)
		if mask[i] == '1' {
			addr |= thisBit
			continue
		}
		// localMask[i] == 'X'

		//     set bit to 1
		addr |= thisBit
		writeMem(addr, val, mask, i+1)
		//     set bit to 0
		addr &= ^thisBit
		writeMem(addr, val, mask, i+1)
		return
	}

	memory[addr] = val
}

func parseMem(line string) {
	rgx, _ := regexp.Compile("^mem\\[([0-9]+)\\] = ([0-9]+)$")
	match := rgx.FindStringSubmatch(line)
	if match == nil { panic("no regex match") }
	strAddr, strVal := match[1], match[2]
	addr, _ := strconv.ParseUint(strAddr, 10, 64)
	val,  _ :=strconv.ParseUint(strVal, 10, 64)
	writeMem(addr, val, mask, 0)
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	memory = make(map[uint64]uint64)

	scanner := bufio.NewScanner(f)
	c := 1
	for scanner.Scan() {
		c++
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			mask =strings.Fields(line)[2]
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
