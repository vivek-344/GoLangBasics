package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	suits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := deck{}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := range d {
		newPosition := r.Intn(len(d))
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

func (d deck) deal(h int) (deck, deck) {
	d1 := d[:h]
	d2 := d[h:]
	return d1, d2
}

func (d deck) saveToFile(name string) error {
	err := os.WriteFile("./output/"+name+".txt", []byte(d.toString()), 0644)
	return err
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func newDeckFromFile(name string) deck {
	data, err := os.ReadFile("./output/" + name + ".txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	cards := deck(strings.Split(string(data), ","))
	return cards
}
