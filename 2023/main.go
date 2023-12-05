package main

import (
	"fmt"

	"adventofcode/adventofcode"
)

func main() {

	scanner, err := adventofcode.FileScanner("adventofcode/day4-input.txt")
	adventofcode.Check(err)
	t := adventofcode.Day4Part2(scanner)
	fmt.Println(t)
}
