package adventofcode

import (
	"bufio"
	"fmt"
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

func Day4Part2(scanner *bufio.Scanner) int {
	total := 0
	lastCard := 0

	copies := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		card := getCardNumber(line)
		winning := getWinningNumbers(line)
		my := getMyNumbers(line)
		matches := intersect.Simple(winning, my)

		lastCard = card

		addCopy(card, copies) // We get one copy for it being in the input

		for cc := 1; cc <= copies[card]; cc++ {
			for i := 1; i <= len(matches); i++ {
				addCopy(card+i, copies)
			}
		}
	}

	for tcard, tcopies := range copies {
		fmt.Printf("%d %d\n", tcard, tcopies)
		if tcard > lastCard {
			continue
		}

		total += tcopies

	}

	return total
}

func addCopy(card int, copies map[int]int) {
	_, ok := copies[card]
	if ok {
		copies[card]++
	} else {
		copies[card] = 1
	}
}

func doubleTimes(x int, times int) int {
	if times == 1 {
		return x
	}

	x = x * 2
	return doubleTimes(x, times-1)
}

func getCardNumber(l string) int {
	i := strings.Split(l, ": ")
	k := strings.Split(i[0], "Card ")

	j, err := strconv.Atoi(strings.Trim(k[1], " "))
	Check(err)

	return j
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
