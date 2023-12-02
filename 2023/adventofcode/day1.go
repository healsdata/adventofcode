package adventofcode

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day1() {
	total := 0

	// scanner := stringScanner(`two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen`)
	scanner, err := FileScanner("adventofcode/day1-input.txt")
	Check(err)

	for scanner.Scan() {
		text := scanner.Text()

		first, last := findFirstAndLastNumber(text)

		line, atoiErr := strconv.Atoi(first + last)
		Check(atoiErr)

		total += line
	}

	fmt.Printf("%d", total)
}

func findFirstAndLastNumber(text string) (string, string) {
	words := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	reLTR := regexp.MustCompile("([0-9]|one|two|three|four|five|six|seven|eight|nine)")
	reRTL := regexp.MustCompile("([0-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)")

	matches := reLTR.FindAllString(text, -1)
	first := matches[0]

	v, ok := words[first]
	if ok {
		first = v
	}

	matches = reRTL.FindAllString(reverse(text), -1)
	last := reverse(matches[0])

	v, ok = words[last]
	if ok {
		last = v
	}

	return first, last
}
