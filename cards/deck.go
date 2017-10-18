package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clovers"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

	// go won't compile code when you declare variable and don't use them, in this instance we are able to use an underscore to replace where the variables for the index would have been because we have no need for them.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// d is a variable that represents the actual deck, and deck on the right represents the type
func (this deck) print() {
	for i, card := range this {
		fmt.Println(i, card)
	}
}

//returning two values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// converts deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// 0666 means that anyone can read or write on this file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	tempDeck := strings.Split(string(bs), ",")
	return deck(tempDeck)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
