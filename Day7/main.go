package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	hands := make(HandCollection, 0)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		cardsString := split[0]
		bet, _ := strconv.Atoi(split[1])
		hand := Hand{
			cards: make([]Card, 0),
			bet:   bet,
		}
		for i := 0; i < len(cardsString); i++ {
			card := newCard(string(cardsString[i]))
			hand.cards = append(hand.cards, card)
		}
		hand.SetValues()
		hands = append(hands, hand)
	}
	sort.Sort(HandCollection(hands))

	for i, hand := range hands {
		fmt.Printf("%d: %v\n", i+1, hand.cards)
		total += hand.bet * (i + 1)
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

type Card struct {
	strValue string
	intValue int
}

func newCard(strValue string) Card {
	intValue := 0
	switch strValue {
	case "A":
		intValue = 14
	case "K":
		intValue = 13
	case "Q":
		intValue = 12
	case "J":
		intValue = 11
	case "T":
		intValue = 10
	default:
		intValue, _ = strconv.Atoi(strValue)
	}
	return Card{
		strValue: strValue,
		intValue: intValue,
	}
}

type Hand struct {
	cards  []Card
	bet    int
	values map[int]int
}

func (h *Hand) SetValues() {
	h.values = make(map[int]int)
	for _, card := range h.cards {
		h.values[card.intValue] += 1
	}
}

func (h Hand) IsFiveOfAKind() bool {
	for _, value := range h.values {
		if value == 5 {
			return true
		}
	}
	return false
}
func (h Hand) IsFourOfAKind() bool {
	for _, value := range h.values {
		if value == 4 {
			return true
		}
	}
	return false
}
func (h Hand) IsFullHouse() bool {
	hasPair, hasThree := false, false
	for _, value := range h.values {
		if value == 2 {
			hasPair = true
		}
		if value == 3 {
			hasThree = true
		}
	}
	return hasPair && hasThree
}
func (h Hand) IsThreeOfAKind() bool {
	for _, value := range h.values {
		if value == 3 {
			return true
		}
	}
	return false
}
func (h Hand) PairCount() int {
	pairCount := 0
	for _, value := range h.values {
		if value == 2 {
			pairCount += 1
		}
	}
	return pairCount
}

type HandCollection []Hand

func (h HandCollection) Len() int {
	return len(h)
}

func (h HandCollection) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h HandCollection) Less(i, j int) bool {
	if h[i].IsFiveOfAKind() != h[j].IsFiveOfAKind() {
		return h[j].IsFiveOfAKind()
	}
	if h[i].IsFiveOfAKind() && h[j].IsFiveOfAKind() {
		return h.LessTieBreak(i, j)
	}
	if h[i].IsFourOfAKind() != h[j].IsFourOfAKind() {
		return h[j].IsFourOfAKind()
	}
	if h[i].IsFourOfAKind() && h[j].IsFourOfAKind() {
		return h.LessTieBreak(i, j)
	}
	if h[i].IsFullHouse() != h[j].IsFullHouse() {
		return h[j].IsFullHouse()
	}
	if h[i].IsFullHouse() && h[j].IsFullHouse() {
		return h.LessTieBreak(i, j)
	}
	if h[i].IsThreeOfAKind() != h[j].IsThreeOfAKind() {
		return h[j].IsThreeOfAKind()
	}
	if h[i].IsThreeOfAKind() && h[j].IsThreeOfAKind() {
		return h.LessTieBreak(i, j)
	}
	if h[i].PairCount() != h[j].PairCount() {
		return h[i].PairCount() < h[j].PairCount()
	}
	return h.LessTieBreak(i, j)
}

func (h HandCollection) LessTieBreak(i, j int) bool {
	iHand := h[i]
	jHand := h[j]
	for k := 0; k < len(iHand.cards); k++ {
		if iHand.cards[k].intValue != jHand.cards[k].intValue {
			return iHand.cards[k].intValue < jHand.cards[k].intValue
		}
	}
	return false
}

type CardCollection []Card

func (h CardCollection) Len() int {
	return len(h)
}

func (h CardCollection) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (c CardCollection) Less(i, j int) bool {
	return c[i].intValue < c[j].intValue
}
