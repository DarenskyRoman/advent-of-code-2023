package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var m = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func keyFindFirst(line string) int {
	for key, value := range m {
		if strings.HasPrefix(line, key) {
			return value
		}
	}
	return keyFindFirst(line[1:])
}

func keyFindLast(line string) int {
	for key, value := range m {
		if strings.HasSuffix(line, key) {
			return value
		}
	}
	return keyFindLast(line[0 : len(line)-1])
}

func main() {

	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	sum1 := 0
	sum2 := 0

	re := regexp.MustCompile("\\d")

	for _, line := range lines {
		digits := re.FindAllString(line, -1)
		num, _ := strconv.Atoi(digits[0] + digits[len(digits)-1])
		sum1 += num

		sum2 += (keyFindFirst(line)*10 + keyFindLast(line))
	}

	fmt.Println("Day 1, part 1: ", sum1)
	fmt.Println("Day 1, part 2: ", sum2)
}
