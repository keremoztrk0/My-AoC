package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	SAFE_MAX_DIFF_VALUE = 3
)

func main() {
	input, err := os.Open("input_day2")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	parsedInput, err := parseInput(input)
	if err != nil {
		panic(err)
	}
	var res1 = part1(parsedInput)
	var res2 = part2(parsedInput)

	fmt.Println("Day 2 result 1:", res1)
	fmt.Println("Day 2 result 2:", res2)
}

const (
	Init       = 0
	Increasing = 1
	Decreasing = 2
)

func part1(input [][]int) int {
	safe_count := 0
	for _, report := range input {
		if isSafeReport(report) {
			safe_count++
		}
	}
	return safe_count
}

func part2(input [][]int) int {
	safe_count := 0
	for _, report := range input {
		if isSafeReport(report) {
			safe_count++
		} else {
			for i := 0; i < len(report); i++ {
				var toleratedReport []int
				toleratedReport = append(toleratedReport, report[:i]...)
				toleratedReport = append(toleratedReport, report[i+1:]...)

				if isSafeReport(toleratedReport) {
					safe_count++
					break
				}
			}
		}
	}
	return safe_count
}

func parseInput(r *os.File) ([][]int, error) {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	var res [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, num := range strings.Split(line, " ") {
			num, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		res = append(res, row)

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func isSafeReport(report []int) bool {
	var diffStatus int = Init
	for col := 0; col < len(report)-1; col++ {
		current := report[col]
		next := report[col+1]
		if next > current {
			if diffStatus == Decreasing {
				return false
			}
			diffStatus = Increasing
		} else if next < current {
			if diffStatus == Increasing {
				return false
			}
			diffStatus = Decreasing

		} else {
			return false
		}
		diff := math.Abs(float64(next - current))

		if diff > SAFE_MAX_DIFF_VALUE {
			return false
		}
	}
	return true
}
