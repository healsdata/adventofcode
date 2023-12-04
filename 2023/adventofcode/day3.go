package adventofcode

import (
	"bufio"
	"strconv"
)

func Day3Part1(scanner *bufio.Scanner) int {
	total := 0

	var lastLine, thisLine, nextLine string

	for scanner.Scan() {
		line := scanner.Text()

		lastLine = thisLine
		thisLine = nextLine
		nextLine = line

		total += findSymbolAdjacentNumbersInThisLine(lastLine, thisLine, nextLine)
	}

	total += findSymbolAdjacentNumbersInThisLine(thisLine, nextLine, "")

	return total
}

func findSymbolAdjacentNumbersInThisLine(lastLine string, thisLine string, nextLine string) int {
	if thisLine == "" {
		return 0
	}

	total := 0

	lr := []rune(lastLine)
	tr := []rune(thisLine)
	nr := []rune(nextLine)

	var curNum string
	var hasSymbol bool

	for i := 0; i < len(tr); i++ {
		if isInt(string(tr[i])) {
			curNum += string(tr[i])

			h := i - 1
			j := i + 1

			if existsAndIsSymbol(tr, h) || existsAndIsSymbol(tr, j) {
				hasSymbol = true
			}

			if existsAndIsSymbol(lr, h) || existsAndIsSymbol(lr, i) || existsAndIsSymbol(lr, j) {
				hasSymbol = true
			}

			if existsAndIsSymbol(nr, h) || existsAndIsSymbol(nr, i) || existsAndIsSymbol(nr, j) {
				hasSymbol = true
			}

			if i+1 < len(tr) {
				continue
			}
		}

		if isInt(curNum) && hasSymbol {
			c, err := strconv.Atoi(curNum)
			Check(err)
			total += c
		}

		curNum = ""
		hasSymbol = false
	}

	return total
}

func existsAndIsSymbol(sr []rune, x int) bool {
	if x >= 0 && x < len(sr) {
		return isSymbol(string(sr[x]))
	}

	return false
}

func isSymbol(s string) bool {
	return !(s == "." || isInt(s))
}
