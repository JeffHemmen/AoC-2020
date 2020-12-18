package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func parseInstruction(line string) (op string, arg int) {
	fields := strings.Fields(line)
	op = fields[0]
	arg, _ = strconv.Atoi(fields[1])
	return
}

func main() {
	acc, ptr := 0, 0
	f, _ := os.Open("input.txt")
	ptrHistory := make(map[int]bool)
	codeOps, codeArgs := make([]string, 0), make([]int, 0)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		op, arg := parseInstruction(scanner.Text())
		codeOps = append(codeOps, op)
		codeArgs = append(codeArgs, arg)
	}

	for i, _ := range codeOps {
		fmt.Println(codeOps[i], codeArgs[i])
	}

	for {
		if ptrHistory[ptr] {
			break
		}
		ptrHistory[ptr] = true
		switch codeOps[ptr] {
			case "acc":
				acc += codeArgs[ptr]
				ptr++
			case "jmp":
				ptr += codeArgs[ptr]
			case "nop":
				ptr++
		}
	}
	fmt.Println(acc)
}
