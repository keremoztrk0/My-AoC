package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	input, err := os.ReadFile("input_day1")
	if err != nil {
		panic(err)
	}
	inputStr := string(input)
	var res1 = part1(inputStr)
	var res2 = part2(inputStr)

	fmt.Println("Day 1 result 1:", res1)
	fmt.Println("Day 1 result 2:", res2)
}

func part1(input string) int {
	left, right := getLeftAndRightSide(input)
	slices.Sort(left)
	slices.Sort(right)
	var res int
	for i := 0; i < len(left); i++ {
		var diff = left[i] - right[i]
		res += int(math.Abs(float64(diff)))
	}
	return res
}

func part2(input string) int {
	left, right := getLeftAndRightSide(input)
	rightGroupMap := createGroupMapForSide(right)
	res := 0
	for i := 0; i < len(left); i++ {
		res += left[i] * rightGroupMap[left[i]]
	}
	return res
}

func getLeftAndRightSide(input string) ([]int, []int) {
	left := []int{}
	right := []int{}
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		sides := strings.Split(line, "   ") // 3 spaces

		left = append(left, stringToInt(sides[0]))
		right = append(right, stringToInt(sides[1]))
	}
	return left, right
}

func stringToInt(input string) int {
	var res int
	fmt.Sscanf(input, "%d", &res)
	return res
}

func createGroupMapForSide(side []int) map[int]int {
	var res = make(map[int]int)
	for _, val := range side {
		res[val]++
	}
	return res
}
