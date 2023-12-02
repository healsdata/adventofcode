package main

import (
	"adventofcode/adventofcode"
	"fmt"
)

func main() {
	//adventofcode.Day1()

	scanner, err := adventofcode.FileScanner("adventofcode/day2-input.txt")
	adventofcode.Check(err)
	//t := adventofcode.Day2Part1(scanner, 12, 13, 14)
	t := adventofcode.Day2Part2(scanner)
	fmt.Println(t)
}
