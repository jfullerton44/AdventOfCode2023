package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		cardTotal := 0
		text := scanner.Text()
		split1 := strings.Split(text, ":")
		nums := strings.Split(split1[1], "|")
		winners := nums[0]
		yourNums := nums[1]
		winNums := strings.Split(strings.TrimSpace(winners), " ")
		winMap := make(map[int]bool)
		for _, winNum := range winNums {
			if winNum != "" {
				intNum, _ := strconv.Atoi(winNum)
				winMap[intNum] = true
			}
		}
		for _, yourNum := range strings.Split(strings.TrimSpace(yourNums), " ") {
			if yourNum != "" {
				intNum, _ := strconv.Atoi(yourNum)
				if _, ok := winMap[intNum]; ok {
					if cardTotal == 0 {
						cardTotal = 1
					} else {
						cardTotal *= 2
					}
				}
			}
		}
		total += cardTotal
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
	cards := make([]int, 200)
	currCard := 0
	for scanner.Scan() {
		cards[currCard] += 1
		cardTotal := 0
		text := scanner.Text()
		split1 := strings.Split(text, ":")
		nums := strings.Split(split1[1], "|")
		winners := nums[0]
		yourNums := nums[1]
		winNums := strings.Split(strings.TrimSpace(winners), " ")
		winMap := make(map[int]bool)
		for _, winNum := range winNums {
			if winNum != "" {
				intNum, _ := strconv.Atoi(winNum)
				winMap[intNum] = true
			}
		}
		for _, yourNum := range strings.Split(strings.TrimSpace(yourNums), " ") {
			if yourNum != "" {
				intNum, _ := strconv.Atoi(yourNum)
				if _, ok := winMap[intNum]; ok {
					if cardTotal == 0 {
						cardTotal = 1
					} else {
						cardTotal += 1
					}
				}
			}
		}
		for i := currCard + 1; i <= currCard+cardTotal; i++ {
			cards[i] += cards[currCard]
		}
		currCard += 1
	}

	for _, card := range cards {
		total += card
	}
	return total, nil
}
