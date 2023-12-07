package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
	}

	return total, nil
}

func round2() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()
	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

	}

	return total, nil
}
