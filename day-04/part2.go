package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"regexp"
)

func  validate_byr(passport map[string]string) bool {
	byr_str := passport["byr"]
	if byr_str == "" {
		return false
	}
	byr, _ := strconv.Atoi(byr_str)
	return (byr >= 1920 && byr <= 2002)
}

func  validate_iyr(passport map[string]string) bool {
	iyr_str := passport["iyr"]
	if iyr_str == "" {
		return false
	}
	iyr, _ := strconv.Atoi(iyr_str)
	return iyr >= 2010 && iyr <= 2020
}

func validate_eyr(passport map[string]string) bool {
	eyr_str := passport["eyr"]
	if eyr_str == "" {
		return false
	}
	eyr, _ := strconv.Atoi(eyr_str)
	return eyr >= 2020 && eyr <= 2030
	return true
}

func validate_hgt(passport map[string]string) bool {
	hgt_str := passport["hgt"]
	if hgt_str == "" {
		return false
	}
	rgx, _ := regexp.Compile("([0-9]+)(cm|in)")
	match := rgx.FindStringSubmatch(hgt_str)
	if match == nil { return false }
	scalar_str, unit := match[1], match[2]
	if unit == "cm" {
		scalar, _ := strconv.Atoi(scalar_str)
		return scalar >= 150 && scalar <= 193
	} else if unit == "in" {
		scalar, _ := strconv.Atoi(scalar_str)
		return scalar >= 59 && scalar <= 76
	}
	return false
}

func validate_hcl(passport map[string]string) bool {
	hcl_str := passport["hcl"]
	if hcl_str == "" {
		return false
	}
	rgx, _ := regexp.Compile("^#[0-9a-f]{6}$")
	match := rgx.FindStringSubmatch(hcl_str)
	return match != nil
}

func validate_ecl(passport map[string]string) bool {
	ecl_str := passport["ecl"]
	if ecl_str == "" {
		return false
	}
	rgx, _ := regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	match := rgx.FindStringSubmatch(ecl_str)
	return match != nil
}

func validate_pid(passport map[string]string) bool {
	pid_str := passport["pid"]
	if pid_str == "" {
		return false
	}
	rgx, _ := regexp.Compile("^[0-9]{9}$")
	match := rgx.FindStringSubmatch(pid_str)
	return match != nil
}

func validate_cid(passport map[string]string) bool {
	return true
}

func validatePassport(passport map[string]string) bool {
	return validate_byr(passport) && validate_iyr(passport) && validate_eyr(passport) && validate_hgt(passport) && validate_hcl(passport) && validate_ecl(passport) && validate_pid(passport) && validate_cid(passport)
}


func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	counter := 0

	var passports []map[string]string

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

	for i, passport_raw := range passports_raw {
		passports = append(passports, make(map[string]string, 0))
		for _, passport_raw_part := range passport_raw {
			field_pairs := strings.Fields(passport_raw_part)

			for _, field_pair := range field_pairs {
				kv := strings.Split(field_pair, ":")
				passports[i][kv[0]] = kv[1]
			}
		}
	}

	for  _, passport := range passports {
		if validatePassport(passport) {
			counter++
		}
	}

	fmt.Println(counter)
}
