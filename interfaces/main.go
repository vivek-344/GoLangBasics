package main

import "fmt"

type bot interface {
	getGreeting() string
}

type hindiBot struct{}

type japaneseBot struct{}

func main() {
	hb := hindiBot{}
	printGreeting(hb)
	jb := japaneseBot{}
	printGreeting(jb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (hindiBot) getGreeting() string {
	return "Namaste!"
}

func (japaneseBot) getGreeting() string {
	return "Konnichiwa!"
}
