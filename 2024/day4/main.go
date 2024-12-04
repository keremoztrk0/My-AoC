package main

import (
	"bufio"
	"os"
	"slices"
)

func main() {
	input := getInput()
	var res1 = part1(input)
	var res2 = part2(input)

	println("Day 3 result 1:", res1)
	println("Day 3 result 2:", res2)
}

func part1(input [][]rune) int {
	return findXMASCount(input)
}

func part2(input [][]rune) int {
	return findXShapedMasesCount(input)
}

const (
	Xmas string = "XMAS"
)

func findXMASCount(input [][]rune) int {
	var res int
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			location := input[y][x]
			if location != 'X' {
				continue
			}
			res += findXmases(input, x, y)

		}
	}
	return res
}

func findXShapedMasesCount(input [][]rune) int {
	var res int
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			location := input[y][x]
			if location != 'A' {
				continue
			}
			if isXShapedMases(input, x, y) {
				res++
			}
		}
	}
	return res
}

func getInput() [][]rune {
	input, err := os.Open("input_day4")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	var res [][]rune
	for scanner.Scan() {
		res = append(res, []rune(scanner.Text()))
	}
	return res
}

func findXmases(input [][]rune, x, y int) int {
	foundCount := 0
	//horizontal
	part := make([]rune, 4)
	if x <= len(input[y])-4 {
		row := input[y]
		if slices.Equal(row[x:x+4], []rune(Xmas)) {
			foundCount++
		}
	}
	//horizontal reverse
	if x >= 3 {
		row := input[y]
		for i := 0; i < 4; i++ {
			part[i] = row[x-i]
		}
		if slices.Equal(part, []rune(Xmas)) {
			foundCount++
		}
		clear(part)
	}

	//vertical
	if y <= len(input)-4 {
		for i := 0; i < 4; i++ {
			part[i] = input[y+i][x]
		}
		if slices.Equal(part, []rune(Xmas)) {
			foundCount++
		}
		clear(part)

	}

	//vertical reverse
	if y >= 3 {
		for i := 0; i < 4; i++ {
			part[i] = input[y-i][x]
		}
		if slices.Equal(part, []rune(Xmas)) {
			foundCount++
		}
		clear(part)
	}
	//diagonal
	if x <= len(input[y])-4 && y <= len(input)-4 {
		for i := 0; i < 4; i++ {
			part[i] = input[y+i][x+i]
		}
		if slices.Equal(part, []rune(Xmas)) {
			foundCount++
		}
		clear(part)
	}

	//diagonal reverse

	if x >= 3 && y >= 3 {
		for i := 0; i < 4; i++ {
			part[i] = input[y-i][x-i]
		}
		if slices.Equal(part, []rune(Xmas)) {
			foundCount++
		}
		clear(part)
	}

	//off diagonal
	if x >= 3 && y <= len(input)-4 {
		for i := 0; i < 4; i++ {
			part[i] = input[y+i][x-i]
		}
		if slices.Equal(part, []rune(Xmas)) {
			foundCount++
		}
		clear(part)
	}

	//off diagonal reverse
	if x <= len(input[y])-4 && y >= 3 {
		for i := 0; i < 4; i++ {
			part[i] = input[y-i][x+i]
		}
		if slices.Equal(part, []rune(Xmas)) {
			foundCount++
		}
		clear(part)
	}
	return foundCount
}

func isXShapedMases(input [][]rune, x, y int) bool {
	if x < 1 || y < 1 || x > len(input[y])-2 || y > len(input)-2 {
		return false
	}
	leftTop := input[y-1][x-1]
	rightTop := input[y-1][x+1]
	leftBottom := input[y+1][x-1]
	rightBottom := input[y+1][x+1]
	if (leftTop == 'M' && rightBottom == 'S' ||
		leftTop == 'S' && rightBottom == 'M') &&
		(leftBottom == 'M' && rightTop == 'S' ||
			leftBottom == 'S' && rightTop == 'M') {
		return true
	}
	return false
}
