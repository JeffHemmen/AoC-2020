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

func evaluatePermutation(codeOps []string, codeArgs []int) (terminates bool, acc int) {
	acc, ptr := 0, 0
	ptrHistory := make(map[int]bool)
	for {
                if ptrHistory[ptr] {
                        terminates = false
			break
                } else if ptr == len(codeOps){
			terminates = true
			return
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
	return
}

func main() {
	f, _ := os.Open("input.txt")
	codeOps, codeArgs := make([]string, 0), make([]int, 0)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		op, arg := parseInstruction(scanner.Text())
		codeOps = append(codeOps, op)
		codeArgs = append(codeArgs, arg)
	}

	modCodeOps := make([]string, len(codeOps))

	for i, op := range codeOps {
		if op == "jmp" {
			copy(modCodeOps, codeOps)
			modCodeOps[i] = "nop"

		} else if op == "nop" {
			copy(modCodeOps, codeOps)
			modCodeOps[i] = "jmp"
		} else {
			continue
		}
		terminates, acc := evaluatePermutation(modCodeOps, codeArgs)
		if terminates {
			fmt.Println(acc)
		}
	}
}
