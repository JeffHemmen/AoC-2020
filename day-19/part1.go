package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"regexp"
)

var rules, parsedRules map[string]string





func parseRule(n string) string {
	if parsedRules[n] != "" {
		return parsedRules[n]
	}
	var res string
	raw := rules[n]

	if strings.Contains(raw, "\"") {
		res = strings.Trim(raw, "\"")
	} else {
		fields := strings.Fields(raw)
		res = "(("
		for _, field := range fields {
			if field == "|" {
				res += ")|("
			} else {
				res += parseRule(field)
			}
		}
		res += "))"
	}

	parsedRules[n] = res
	return res
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	rules = make(map[string]string)
	parsedRules = make(map[string]string)

	scanner.Scan()
	line := scanner.Text()
	for line != "" && scanner.Scan() {
		// read rules
		kv := strings.Split(line, ": ")
		rules[kv[0]] = kv[1]
		line = scanner.Text()
	}


	regexraw := parseRule("0")
	regexraw = "^" + regexraw + "$"
	// fmt.Println(regex)
	rgx, _ := regexp.Compile(regexraw)
	counter := 0

	for scanner.Scan() {
		// read messages
		line = scanner.Text()
		match := rgx.FindStringSubmatch(line)
		if match != nil {
			counter++
		}
	}

	fmt.Println(counter)
}
