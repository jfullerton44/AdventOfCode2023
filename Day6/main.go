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
	total := 1
	scanner := bufio.NewScanner(f)
	races := make([]Race, 0)
	for scanner.Scan() {
		line := scanner.Text()
		times := strings.Split(line, " ")
		for _, time := range times {
			timeInt, err := strconv.Atoi(time)
			if err == nil {
				race := Race{
					time: timeInt,
				}
				races = append(races, race)
			}
		}
		scanner.Scan()
		line = scanner.Text()
		raceCount := 0
		distances := strings.Split(line, " ")
		for _, distance := range distances {
			distanceInt, err := strconv.Atoi(distance)
			if err == nil {
				races[raceCount].distance = distanceInt
				raceCount += 1
			}
		}
	}
	results := make([]int, len(races))
	for i, race := range races {
		currWays := 0
		for holdTime := 0; holdTime < race.time; holdTime++ {
			if holdTime*(race.time-holdTime) > race.distance {
				currWays += 1
			}
		}
		results[i] = currWays
		total *= currWays
	}
	return total, nil
}

func round2() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()
	total := 1
	scanner := bufio.NewScanner(f)
	races := make([]Race, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1)
		times := strings.Split(line, ":")
		for _, time := range times {
			timeInt, err := strconv.Atoi(time)
			if err == nil {
				race := Race{
					time: timeInt,
				}
				races = append(races, race)
			}
		}
		scanner.Scan()
		line = scanner.Text()
		raceCount := 0
		line = strings.Replace(line, " ", "", -1)
		distances := strings.Split(line, ":")
		for _, distance := range distances {
			distanceInt, err := strconv.Atoi(distance)
			if err == nil {
				races[raceCount].distance = distanceInt
			}
		}
	}
	results := make([]int, len(races))
	for i, race := range races {
		currWays := 0
		for holdTime := 0; holdTime < race.time; holdTime++ {
			if holdTime*(race.time-holdTime) > race.distance {
				currWays += 1
			}
		}
		results[i] = currWays
		total *= currWays
	}
	return total, nil
}

type Race struct {
	time     int
	distance int
}
