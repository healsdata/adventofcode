package adventofcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay4(t *testing.T) {
	t.Run("acceptance1", func(t *testing.T) {
		ex := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

		scanner, err := stringScanner(ex)
		require.NoError(t, err)

		//assert.Equal(t, 13, Day4Part1(scanner))
		assert.Equal(t, 30, Day4Part2(scanner))
	})

}

func TestDay4getWinningNumbers(t *testing.T) {

	type test struct {
		line     string
		expected []int
	}

	tests := []test{
		{line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", expected: []int{41, 48, 83, 86, 17}},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, getWinningNumbers(tc.line))
		})
	}
}

func TestDay4getMyNumbers(t *testing.T) {

	type test struct {
		line     string
		expected []int
	}

	tests := []test{
		{line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", expected: []int{83, 86, 6, 31, 17, 9, 48, 53}},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, getMyNumbers(tc.line))
		})
	}
}
