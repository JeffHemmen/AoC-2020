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

	for ia, a := range list {
		for ib, b := range list[ia+1:] {
			for _, c := range list[ia+ib+1:] {
				if  a + b + c == 2020 {
					fmt.Println(a * b * c)
					return
				}
			}
		}
	}
}
