package adventofcode

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestDay2(t *testing.T) {

	t.Run("acceptance1", func(t *testing.T) {
		ex := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

		scanner, err := stringScanner(ex)
		require.NoError(t, err)

		assert.Equal(t, 8, Day2(scanner, 12, 13, 14))

	})

}

func TestDay2getMaxColorsFromGame(t *testing.T) {

	type test struct {
		input   string
		expectR int
		expectG int
		expectB int
	}

	tests := []test{
		{input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", expectR: 4, expectG: 2, expectB: 6},
		{input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", expectR: 20, expectG: 13, expectB: 6},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r, g, b := getMaxColorsFromGame(tc.input)
			assert.Equal(t, tc.expectR, r)
			assert.Equal(t, tc.expectG, g)
			assert.Equal(t, tc.expectB, b)
		})
	}
}
