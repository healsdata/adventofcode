package adventofcode

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func Day2(scanner *bufio.Scanner, r int, g int, b int) int {

	total := 0

	for scanner.Scan() {
		game := scanner.Text()

		gr, gg, gb := getMaxColorsFromGame(game)

		if gr <= r && gg <= g && gb <= b {
			reg := regexp.MustCompile("^Game (?P<I>[0-9]+):")
			m := reg.FindStringSubmatch(game)
			i, err := strconv.Atoi(m[len(m)-1])
			Check(err)
			total += i
		}

	}

	return total
}

func getMaxColorsFromGame(game string) (int, int, int) {
	var r, g, b int

	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	i := strings.Split(game, ":")
	j := strings.Trim(i[1], " ")
	pulls := strings.Split(j, ";")

	for _, pull := range pulls {
		pr := getColorFromPull(pull, "red")
		if pr >= r {
			r = pr
		}

		pg := getColorFromPull(pull, "green")
		if pg >= g {
			g = pg
		}

		pb := getColorFromPull(pull, "blue")
		if pb >= b {
			b = pb
		}
	}

	return r, g, b
}

func getColorFromPull(pull string, color string) int {
	r := regexp.MustCompile("( |^)(?P<I>[0-9]+) " + color)
	m := r.FindStringSubmatch(pull)
	if m == nil {
		return 0
	}

	i, err := strconv.Atoi(m[len(m)-1])
	Check(err)

	return i
}
