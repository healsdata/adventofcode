package adventofcode

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
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

func FileScanner(fn string) (*bufio.Scanner, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(bufio.NewReader(f)), nil
}

func stringScanner(s string) (*bufio.Scanner, error) {
	return bufio.NewScanner(strings.NewReader(s)), nil
}

func isInt(s string) bool {
	if s == "" {
		return false
	}

	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return true
}

func ChunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}
