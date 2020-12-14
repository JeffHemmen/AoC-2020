package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("part1.txt")
	defer f.Close()

	var list []int // = make([]int, 256) // starting size, will expand as needed

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		list = append(list, n)
	}
	//fmt.Println(list)

	for i, a := range list {
		for _, b := range list[i+1:] {
			if  a + b == 2020 {
				fmt.Println(a * b)
				return
			}
		}
	}
}
