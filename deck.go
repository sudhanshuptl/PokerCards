package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create new type deck which is slice of string
type deck []string

func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") // converting deck type to slice of string then joining comman seperated
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is permissiom granted to drive
}

func (d deck) suffle() {
	/* suffle each position value with random generate dposition value */

	source := rand.NewSource(time.Now().UnixNano()) // time.Now().UnixNano() act as seed value for randaon number generator
	rand.New(source)
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
func getDeckFromFile(filename string) deck {
	resultbytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error occured ->", err)
		os.Exit(1) //exit with error
	}
	var d deck
	for _, v := range strings.Split(string(resultbytes), ",") {
		d = append(d, v)
	}
	return d
}

func newDeck() deck {
	// return newly created fresh deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
