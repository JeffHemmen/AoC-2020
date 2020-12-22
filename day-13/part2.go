package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"sort"
)

func primeFactors(a int) []int {
	res := make([]int, 0)
	for i := 2; a != 1; {
		if a % i == 0 {
			res = append(res, i)
			a = a / i
		} else {
			i++
		}
	}
	return res
}

func lcm(a, b int) int {
	pfa, pfb := primeFactors(a), primeFactors(b)
	pfc := make([]int, 0)
	for len(pfa) > 0 && len(pfb) > 0 {
		if pfa[0] == pfb[0] {
			pfc = append(pfc, pfa[0])
			pfa, pfb = pfa[1:], pfb[1:]
		} else if  pfa[0] < pfb[0] {
			pfc = append(pfc, pfa[0])
			pfa = pfa[1:]
		} else {
			pfc = append(pfc, pfb[0])
			pfb = pfb[1:]
		}
	}
	pfc = append(pfc, pfa...)
	pfc = append(pfc, pfb...)
	res := 1
	for _, f := range pfc {
		res *= f
	}
	return res
}

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	scanner.Scan() // ignore first row
	scanner.Scan()
	f.Close()
	schedule := scanner.Text()
	xbuses := strings.Split(schedule, ",")

	type subsequentBuses struct {
		bus int
		interval int
	}

	firstBus := 0
	buses := make([]subsequentBuses, 0)

	acc := 1
	for _, bus := range xbuses {
		if bus == "x" {
			acc++
			continue
		}
		intbus, _ := strconv.Atoi(bus)
		if firstBus == 0 {
			firstBus = intbus // never "x"
			continue
		}
		buses = append(buses, subsequentBuses{intbus, acc})
		acc++
	}

	sort.Slice(buses, func (i, j int) bool { return buses[i].bus < buses[j].bus })
	// sorting low->high  results in 487 iterations overall
	// not sorting at all results in 515 iterations overall
	// sorting ligh->low  results in 849 iterations overall

	// bigO:=0
	for factor, increment, nextIncidence := 1, 1, 0; ; factor += increment {
		time := firstBus * factor
		// fmt.Println(firstBus, factor, time)
		if (time + buses[nextIncidence].interval) % buses[nextIncidence].bus == 0 {
			if nextIncidence == len(buses) - 1 {
				fmt.Println(time)
				break
			}
			increment = lcm(increment, buses[nextIncidence].bus)
			nextIncidence++
		}
		// bigO++
	}

	// fmt.Println(bigO)

}
