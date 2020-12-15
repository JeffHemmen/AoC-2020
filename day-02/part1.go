package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
)

func checkPass(pass string, char rune, min, max int) bool {
	count := 0
	//fmt.Println(pass, char, min, max)
	for _, r := range pass {
		//fmt.Println("    ", r)
		if r == char {
			count++
		}
	}
	//fmt.Println(count)
	return  count >= min && count <= max
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0

	for scanner.Scan() {
		rawline := scanner.Text()
		line := strings.Fields(rawline)
		var char rune
		//fmt.Println(line, line[0])
		for _, char = range line[1] {break}
		limits := strings.FieldsFunc(line[0], func(s rune ) bool { return s == '-'})
		min, _ := strconv.Atoi(limits[0])
		max, _ := strconv.Atoi(limits[1])
		if checkPass(line[2], char, min, max) {
			count++
		}
	}
	fmt.Println(count)
}
