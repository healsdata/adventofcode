package adventofcode

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestDay3(t *testing.T) {
	t.Run("acceptance1", func(t *testing.T) {
		ex := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

		scanner, err := stringScanner(ex)
		require.NoError(t, err)

		assert.Equal(t, 4361, Day3Part1(scanner))
	})
}

func TestDay3findSymbolAdjacentNumbersInThisLine(t *testing.T) {

	type test struct {
		lastLine string
		thisLine string
		nextLine string
		expected int
	}

	tests := []test{
		{lastLine: "", thisLine: "", nextLine: "467..114..", expected: 0},
		{lastLine: "", thisLine: "467..114..", nextLine: "...*......", expected: 467},
		{lastLine: "...$.*....", thisLine: ".664.598..", nextLine: "", expected: 664 + 598},
		{lastLine: "", thisLine: "617*......", nextLine: "", expected: 617},
		{lastLine: "...*......", thisLine: "617.......", nextLine: "", expected: 617},
		{lastLine: "..*.......", thisLine: "617.......", nextLine: "", expected: 617},
		{lastLine: "*.........", thisLine: "617.......", nextLine: "", expected: 617},
		{lastLine: ".........*", thisLine: ".......617", nextLine: "", expected: 617},
		{lastLine: "......*...", thisLine: ".......617", nextLine: "", expected: 617},
		{lastLine: ".......*..", thisLine: ".......617", nextLine: "", expected: 617},
		{
			lastLine: "................................405*.......................29...%1...................*........................754...#.........*.............",
			thisLine: "961.........396.....................472.......225..739..............415............451......................................904.............",
			nextLine: "",
			expected: 472 + 451 + 904,
		},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, findSymbolAdjacentNumbersInThisLine(tc.lastLine, tc.thisLine, tc.nextLine))
		})
	}
}
