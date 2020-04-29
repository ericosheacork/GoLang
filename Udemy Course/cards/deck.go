// creating a deck object here
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "King", "Queen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// take a deck type convert to a string and then to a byte slice for saving (use a helper function)
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveDeck(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func readDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR OCCURED")
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

//this piece of code "Randomizes" but in the same way every time so not random
func (d deck) shuffle() {
	// the folloeing 2 lines allow for better randomisation
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		//this is a psuedo random number generator, if we did not use r (the same seed is always used)
		newPos := r.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], d[index]
	}
}
