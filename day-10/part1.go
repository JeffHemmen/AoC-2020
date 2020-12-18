package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"sort"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	adapters := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		this, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, this)
	}
	sort.Ints(adapters)
	differences := make(map[int]int)
	previous := 0
	for _, adapter := range adapters {
		differences[adapter-previous]++
		previous = adapter
	}
	differences[3]++
	fmt.Println(differences[1] * differences[3])
}
