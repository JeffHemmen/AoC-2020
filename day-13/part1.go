package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	est, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	schedule := scanner.Text()
	xbuses := strings.Split(schedule, ",")
	buses := make([]int, 0)
	for _, bus := range xbuses {
		if bus == "x" { continue }
		intbus, _ := strconv.Atoi(bus)
		buses = append(buses, intbus)
	}

	nextBuses := make([]int, 0)
	for id, bus := range buses {
		nextBuses = append(nextBuses, 0)
		for i := 0; nextBuses[id] < est; i++ {
			nextBuses[id] = bus * i 
		}
	}

	var lowestID int = 0

	for id, bus := range nextBuses {
		if bus < nextBuses[lowestID] {
			lowestID = id
		}
	}
		
	fmt.Println(buses[lowestID] * (nextBuses[lowestID] - est))
}
