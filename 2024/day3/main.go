package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := getInput()
	var res1 = part1(input)
	var res2 = part2(input)

	fmt.Println("Day 1 result 1:", res1)
	fmt.Println("Day 1 result 2:", res2)
}

func part1(input string) int {
	muls := getMuls(input)
	var res int
	for _, mul := range muls {
		res += calculateMul(mul)
	}
	return res
}

func part2(input string) int {
	muls := getValidMuls(input)
	var res int
	for _, mul := range muls {
		res += calculateMul(mul)
	}
	return res
}

func getMuls(input string) []string {
	mullRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return mullRegex.FindAllString(input, -1)
}

func getValidMuls(input string) []string {
	excludeOpRegex := regexp.MustCompile(`(?s)don't\(\).*?($|do\(\))`)
	input = excludeOpRegex.ReplaceAllString(input, "")
	return getMuls(input)
}
func calculateMul(mul string) int {
	numberRegex := regexp.MustCompile(`\d{1,3}`)
	numbers := numberRegex.FindAllString(mul, -1)
	var res int = 1
	for _, number := range numbers {
		number, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		res *= number
	}
	return res
}

func getInput() string {
	input, err := os.ReadFile("input_day3")
	if err != nil {
		panic(err)
	}

	return string(input)
}
