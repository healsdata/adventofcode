package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	/*
	   	ex := `two1nine
	   eightwothree
	   abcone2threexyz
	   xtwone3four
	   4nineeightseven2
	   zoneight234
	   7pqrstsixteen`
	*/
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

	f, err := os.Open("day1-input.txt")
	check(err)

	reLTR := regexp.MustCompile("([0-9]|one|two|three|four|five|six|seven|eight|nine)")
	reRTL := regexp.MustCompile("([0-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)")

	total := 0

	//scanner := bufio.NewScanner(strings.NewReader(ex))

	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		text := scanner.Text()

		fmt.Println(text)

		matches := reLTR.FindAllString(text, -1)
		first := matches[0]

		fmt.Println(matches)
		v, ok := words[first]
		if ok {
			first = v
		}

		matches = reRTL.FindAllString(reverse(text), -1)
		last := reverse(matches[0])
		fmt.Println(matches)

		v, ok = words[last]
		if ok {
			last = v
		}

		line := first + last
		fmt.Println(line)

		li, err := strconv.Atoi(line)
		check(err)

		total += li
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	fmt.Printf(strconv.Itoa(total))

}
