package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var in *bufio.Reader
var out *bufio.Writer
var cardPoints = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}
var values = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var spades = [4]string{"C", "D", "H", "S"}

type card struct {
	name  string
	point int
	value string
	suit  string
}

func newCard(name string) *card {
	value := string(name[0])
	point := cardPoints[value]
	suit := string(name[1])
	return &card{
		name:  name,
		point: point,
		value: value,
		suit:  suit,
	}
}

func (c *card) bestCard(c2 *card) card {
	if c.point >= c2.point {
		return *c
	}
	return *c2
}

type hand struct {
	card1      card
	card2      card
	totalPoint int
}

func newHand(card1, card2 card) *hand {
	var totalPoint int
	if card1.point > card2.point {
		totalPoint = card1.point
	} else if card1.point < card2.point {
		totalPoint = card2.point
	} else {
		totalPoint = card1.point + card2.point
	}
	return &hand{
		card1:      card1,
		card2:      card2,
		totalPoint: totalPoint,
	}
}

func bestHand(hands []hand) hand {
	best := 0
	for i, h := range hands {
		if h.totalPoint > hands[best].totalPoint {
			best = i
		}
	}
	return hands[best]
}

func uselessOpponentsCards(hands []hand) []string {
	usedCards := map[string]bool{}
	for _, value := range values {
		for _, spade := range spades {
			usedCards[value+spade] = false
		}
	}
	for _, h := range hands {
		if h.card1.point == h.card2.point {
			for _, spade := range spades {
				usedCards[h.card1.value+spade] = true
			}
		} else {
			usedCards[h.card1.name] = true
			usedCards[h.card2.name] = true
		}
	}
	rez := make([]string, 0)
	for k, used := range usedCards {
		if !used {
			rez = append(rez, k)
		}
	}
	slices.Sort(rez)
	return rez
}

func (c *card) remainingSuits(hands []hand) []string {
	usedSuits := map[string]bool{
		"C": false,
		"D": false,
		"H": false,
		"S": false,
	}
	value := c.point
	for _, h := range hands {
		if h.card1.point == value {
			usedSuits[h.card1.suit] = true
		}
		if h.card2.point == value {
			usedSuits[h.card2.suit] = true
		}
	}
	rez := make([]string, 0)
	for suit, used := range usedSuits {
		if !used {
			rez = append(rez, suit)
		}
	}
	return rez
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		processTest()
	}
}

func processTest() {
	var n int
	fmt.Fscan(in, &n)
	data := make([][2]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &data[i][0], &data[i][1])
	}
	poker(n, data)
}

func poker(n int, data [][2]string) {
	hands := make([]hand, n)
	for i := 0; i < n; i++ {
		c1 := *newCard(data[i][0])
		c2 := *newCard(data[i][1])
		hands[i] = *newHand(c1, c2)
	}
	myHand := hands[0]
	myBestCard := myHand.card1
	if myHand.card1.point < myHand.card2.point {
		myBestCard = myHand.card2
	}
	bestHand := bestHand(hands)

	remainingSuits := myBestCard.remainingSuits(hands)
	if myHand.totalPoint == bestHand.totalPoint {
		if len(remainingSuits) == 0 {
			rest := uselessOpponentsCards(hands)
			fmt.Fprintln(out, len(rest))
			fmt.Fprintln(out, strings.Join(rest, "\n"))
		} else {
			for _, suit := range remainingSuits {
				fmt.Fprintln(out, len(remainingSuits))
				fmt.Fprintln(out, myBestCard.name+suit)
			}
		}
	} else {
		if len(remainingSuits) == 0 {
			fmt.Fprintln(out, 0)
		} else {
			for _, suit := range remainingSuits {
				fmt.Fprintln(out, len(remainingSuits))
				fmt.Fprintln(out, myBestCard.name+suit)
			}
		}
	}
}
