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
	scanner := bufio.NewScanner(f)
	transforms := make([][]Transform, 0, 10)
	currIndex := 0
	scanner.Scan()
	seeds := scanner.Text()
	seedsSide := strings.Split(seeds, ":")
	seedsSplit := strings.Split(strings.TrimSpace(seedsSide[1]), " ")
	seedVals := make([]int, 0, len(seedsSplit))
	for _, seed := range seedsSplit {
		if seed != "" {
			seedVal, _ := strconv.Atoi(seed)
			seedVals = append(seedVals, seedVal)
		}
	}
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		currTransform := make([]Transform, 0)
		scanner.Scan()
		for scanner.Text() != "" {
			text := scanner.Text()
			currIndex += 1
			line := strings.Split(text, " ")
			destination, _ := strconv.Atoi(line[0])
			source, _ := strconv.Atoi(line[1])
			count, _ := strconv.Atoi(line[2])
			transfrom := Transform{
				destination: destination,
				source:      source,
				count:       count,
			}
			scanner.Scan()
			currTransform = append(currTransform, transfrom)
		}
		transforms = append(transforms, currTransform)
	}

	for _, transform := range transforms {
		for i, seedVal := range seedVals {
			for _, currTransform := range transform {
				if seedVal > currTransform.source && seedVal <= currTransform.source+currTransform.count {
					seedVal = currTransform.destination + (seedVal - currTransform.source)
					break
				}
			}
			seedVals[i] = seedVal
		}
	}
	min := seedVals[0]
	for _, seedVal := range seedVals {
		if seedVal < min {
			min = seedVal
		}
	}
	return min, nil
}

func round2() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	transforms := make([][]Transform, 0, 10)
	currIndex := 0
	scanner.Scan()
	seeds := scanner.Text()
	seedsSide := strings.Split(seeds, ":")
	seedsSplit := strings.Split(strings.TrimSpace(seedsSide[1]), " ")
	seedRanges := make([]SeedRange, 0, len(seedsSplit)/2)
	for i, seed := range seedsSplit {
		if i%2 == 0 {
			if seed != "" {
				seedVal, _ := strconv.Atoi(seed)
				length, _ := strconv.Atoi(seedsSplit[i+1])
				seedRange := SeedRange{
					start: seedVal,
					count: length,
				}
				seedRanges = append(seedRanges, seedRange)
				i += 1
			}
		}

	}
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		currTransform := make([]Transform, 0)
		scanner.Scan()
		for scanner.Text() != "" {
			text := scanner.Text()
			currIndex += 1
			line := strings.Split(text, " ")
			destination, _ := strconv.Atoi(line[0])
			source, _ := strconv.Atoi(line[1])
			count, _ := strconv.Atoi(line[2])
			transfrom := Transform{
				destination: destination,
				source:      source,
				count:       count,
			}
			scanner.Scan()
			currTransform = append(currTransform, transfrom)
		}
		transforms = append(transforms, currTransform)
	}

	for _, transform := range transforms {
		for i := 0; i < len(seedRanges); i++ {
			seedRange := seedRanges[i]
			fmt.Printf("%d: %d ->", seedRange.start, seedRange.count)
			for _, currTransform := range transform {
				if seedRange.start >= currTransform.source && seedRange.start < currTransform.source+currTransform.count {
					rangeSize := (currTransform.source + currTransform.count) - seedRange.start
					if rangeSize > seedRange.count {
						rangeSize = seedRange.count
					}
					if rangeSize < seedRange.count {
						newSource := SeedRange{
							start: seedRange.start + rangeSize,
							count: seedRange.count - rangeSize,
						}
						seedRanges = append(seedRanges, newSource)
					}
					seedRange.start = currTransform.destination + (seedRange.start - currTransform.source)
					seedRange.count = rangeSize
					break
				} else if currTransform.source+currTransform.count > seedRange.start && seedRange.start > currTransform.source {
					rangeSize := currTransform.source + currTransform.count - seedRange.start
					if rangeSize > seedRange.count {
						rangeSize = seedRange.count
					}
					if rangeSize < seedRange.count {
						newSource := SeedRange{
							start: seedRange.start,
							count: seedRange.count - rangeSize,
						}
						seedRanges = append(seedRanges, newSource)
					}
					seedRange.start = currTransform.destination + (seedRange.start - currTransform.source)
					seedRange.count = rangeSize
					break
				}
			}
			fmt.Printf("%d: %d\n", seedRange.start, seedRange.count)
			seedRanges[i] = seedRange
		}
		fmt.Println()
	}
	min := seedRanges[0].start
	for _, seedVal := range seedRanges {
		if seedVal.start < min {
			min = seedVal.start
		}
	}
	return min, nil
}

type Transform struct {
	destination int
	source      int
	count       int
}

type SeedRange struct {
	start int
	count int
}
