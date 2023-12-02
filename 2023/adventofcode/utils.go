package adventofcode

import (
	"bufio"
	"os"
	"strings"
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

func fileScanner(fn string) (*bufio.Scanner, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(bufio.NewReader(f)), nil
}

func stringScanner(s string) (*bufio.Scanner, error) {
	return bufio.NewScanner(strings.NewReader(s)), nil
}
