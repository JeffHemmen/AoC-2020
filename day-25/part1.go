package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

const subjNum = 7
const divisor = 20201227

func transform(value, subjNum int) int {
	return value * subjNum % divisor
}

func main() {

	labels := [2]string{"Card", "Door"}
	var pubKeys []int     = make([]int, 2)
	var loopsizes []int   = make([]int, 2)
	// var subjectNums []int = make([]int, 2)


	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	pubCardStr := scanner.Text()
	scanner.Scan()
	pubDoorStr := scanner.Text()
	f.Close()
	pubKeys[0], _ = strconv.Atoi(pubCardStr)
	pubKeys[1], _ = strconv.Atoi(pubDoorStr)

	// find loop size
	value := 1
	foundId := -1
	for i := 1; foundId == -1 ; i++ {
		value = transform(value, subjNum)
		for id := 0; id <= 1; id++ {
			if value == pubKeys[id] {
				foundId = id
				loopsizes[id] = i
				break
			}
		}
	}

	debug := false
	if debug {
		// if I just comment the line below out, Go complains about unused var
		fmt.Println(labels[foundId], "'s loopsize:", loopsizes[foundId])
	}

	value = 1
	for i:= 0; i < loopsizes[foundId]; i++ {
		value = transform(value, pubKeys[1-foundId])
	}
	fmt.Println(value)

}
