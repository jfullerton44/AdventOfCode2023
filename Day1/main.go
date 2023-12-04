package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	rd1, err := round1()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 1: %d\n", rd1)

	rd2, err := round2()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 2: %d\n", rd2)

}

func round1() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		var first, last *int
		for _, c := range text {
			if num, err := strconv.Atoi(string(c)); err == nil {
				if first == nil {
					first = &num
					last = &num
				} else {
					last = &num
				}
			}
		}
		total += *first*10 + *last
	}

	return total, nil
}

func round2() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		firstSet := false
		var first, last int
		for i, c := range text {
			if num, err := strconv.Atoi(string(c)); err == nil {
				if !firstSet {
					first = num
					last = num
					firstSet = true
				} else {
					last = num
				}
			} else {
				if substr(text, i, 3) == "one" {
					if !firstSet {
						first = 1
						last = 1
						firstSet = true
					} else {
						last = 1
					}
				}
				if substr(text, i, 3) == "two" {
					if !firstSet {
						first = 2
						last = 2
						firstSet = true
					} else {
						last = 2
					}
				}
				if substr(text, i, 5) == "three" {
					if !firstSet {
						first = 3
						last = 3
						firstSet = true
					} else {
						last = 3
					}
				}
				if substr(text, i, 4) == "four" {
					if !firstSet {
						first = 4
						last = 4
						firstSet = true
					} else {
						last = 4
					}
				}
				if substr(text, i, 4) == "five" {
					if !firstSet {
						first = 5
						last = 5
						firstSet = true
					} else {
						last = 5
					}
				}
				if substr(text, i, 3) == "six" {
					if !firstSet {
						first = 6
						last = 6
						firstSet = true
					} else {
						last = 6
					}
				}
				if substr(text, i, 5) == "seven" {
					if !firstSet {
						first = 7
						last = 7
						firstSet = true
					} else {
						last = 7
					}
				}
				if substr(text, i, 5) == "eight" {
					if !firstSet {
						first = 8
						last = 8
						firstSet = true
					} else {
						last = 8
					}
				}
				if substr(text, i, 4) == "nine" {
					if !firstSet {
						first = 9
						last = 9
						firstSet = true
					} else {
						last = 9
					}
				}
				if substr(text, i, 4) == "zero" {
					if !firstSet {
						first = 0
						last = 0
						firstSet = true
					} else {
						last = 0
					}
				}
			}

		}
		total += first*10 + last
	}

	return total, nil
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
