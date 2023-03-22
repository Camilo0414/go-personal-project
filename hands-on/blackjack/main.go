package main

import (
	"fmt"
	"strconv"
)

type deck map[string]string

func (d deck) calculateValue(k string) string {
	return d[k]
}
func (d deck) sumValues(v1, v2 string) int {
	i1, _ := strconv.Atoi(d.calculateValue(v1))
	i2, _ := strconv.Atoi(d.calculateValue(v2))
	return i1 + i2
}

func winDeal(dc string) bool {
	//returns true if the card of the dealer is a winner card. otherwise returns false

	winHand := []string{"ace", "jack", "queen", "king", "ten"}
	var b = false
	for _, val := range winHand {
		if dc != val && b != true {
			b = false
		} else {
			b = true
		}
	}
	fmt.Println(b, "windeal")
	return b
}

func (d deck) highDeal(dc string) bool {
	highHand := []string{"seven", "eight", "nine", "ten", "jack", "queen", "king", "ace"}
	var b = false
	for _, val := range highHand {
		if dc != val && b != true {
			b = false
		} else {
			b = true
		}
	}
	return b
}

func (d deck) caseOne(card1, card2 string) bool {
	return card1 == "ace" && card2 == "ace"
}

func (d deck) caseTwo(c1, c2, dc string) bool {
	val := d.sumValues(c1, c2)
	var b = true
	fmt.Println("Case two:", c1, c2)
	fmt.Println("Case two:", val)
	if val == 21 && winDeal(dc) != true {
		fmt.Println(winDeal(dc), "inside if")
		b = true
	} else {
		b = false
	}
	fmt.Println(b)
	return b
}

func (d deck) caseThree(c1, c2, dc string) bool {
	var b = true
	if d.sumValues(c1, c2) >= 17 && d.sumValues(c1, c2) <= 20 {
		b = true
	} else {
		//if d.highDeal(dc) == true {
		b = false
	}
	return b
}

func (d deck) firstTurn(card1, card2, dealerCard string) {
	fmt.Println("firstTrun:", card1, card2, dealerCard)
	switch {
	case d.caseOne(card1, card2) == true:
		fmt.Println("P")
	case d.caseTwo(card1, card2, dealerCard) == true && d.sumValues(card1, card2) == 21:
		fmt.Println("W")
	case d.caseTwo(card1, card2, dealerCard) == false && d.sumValues(card1, card2) == 21:
		fmt.Println("S-this")
	case d.caseThree(card1, card2, dealerCard) == true && d.sumValues(card1, card2) >= 12:
		fmt.Println("S-that")
	case d.caseThree(card1, card2, dealerCard) == false && d.sumValues(card1, card2) >= 12 && d.highDeal(dealerCard) == true:
		fmt.Println("H")
	case d.caseThree(card1, card2, dealerCard) == false && d.sumValues(card1, card2) >= 12 && d.highDeal(dealerCard) == false:
		fmt.Println("S")
	case d.sumValues(card1, card2) <= 11:
		fmt.Println("H")
	default:
		fmt.Println("Learn to play pal")
	}
}

func main() {
	d := deck{
		"ace":   "11",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"ten":   "10",
		"jack":  "10",
		"queen": "10",
		"king":  "10",
		"other": "0",
	}
	d.firstTurn("seven", "four", "other")
}
