package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	counter := 0

	var passports_raw [][]string
	passports_raw = append(passports_raw, make([]string, 0))
	passport_idx := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passport_idx++
			passports_raw = append(passports_raw, make([]string, 0))
			continue
		}
		passports_raw[passport_idx] = append(passports_raw[passport_idx], scanner.Text())
	}

	for _, passport_raw := range passports_raw {
		field_keys := make([]string, 0)
		for _, passport_raw_part := range passport_raw {
			field_pairs := strings.Fields(passport_raw_part)

			for _, field_pair := range field_pairs {
				field_keys = append(field_keys, strings.Split(field_pair, ":")[0])
			}
		}
		if len(field_keys) == 8 {
			counter++
			fmt.Println("Valid Passport")
			continue
		}
		if len(field_keys) <= 6 {
			fmt.Println("Only %n fields", len(field_keys))
			continue
		}
		fmt.Println(field_keys)

		npc := true
		for _, key := range field_keys {
			if key == "cid" {
				npc = false
				break
			}
		}
		if npc {
			counter++
			fmt.Println("North Pole Credentials")
		}
	}
	fmt.Println(counter)
}
