package adventofcode

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/juliangruber/go-intersect"
)

func Day4Part1(scanner *bufio.Scanner) int {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		winning := getWinningNumbers(line)
		my := getMyNumbers(line)
		matches := intersect.Simple(winning, my)
		if len(matches) > 0 {
			total += doubleTimes(1, len(matches))
		}
	}

	return total
}

func doubleTimes(x int, times int) int {
	if times == 1 {
		return x
	}

	x = x * 2
	return doubleTimes(x, times-1)
}

func getWinningNumbers(l string) []int {
	i := strings.Split(l, ": ")
	k := strings.Split(i[1], "|")

	return chunkStringToInt(k[0])
}

func getMyNumbers(l string) []int {
	i := strings.Split(l, ": ")
	k := strings.Split(i[1], "|")
	return chunkStringToInt(k[1])
}

func chunkStringToInt(j string) []int {
	wStr := ChunkString(j, 3)
	var wInt []int
	for _, v := range wStr {
		vI, err := strconv.Atoi(strings.Trim(v, " "))
		Check(err)
		wInt = append(wInt, vI)
	}

	return wInt
}
