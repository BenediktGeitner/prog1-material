package main

import (
	"fmt"
	"math/rand"
)

// Readnumber: frag den Benutzer nach irgendeiner Zahl und liefert diese zurück
func ReadNumber() int {
	fmt.Print("Rate eine Zahl: ")
	var n int
	fmt.Scan(&n)
	return n
}

func GuessingGame() {
	my_number := rand.Intn(101) - 50
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if NumberGood(guess, my_number) {
			fmt.Println("Richtig geraten :-)")
			return
		} else if guess < my_number {
			fmt.Println("Meine Zahl ist größer")
		} else {
			fmt.Println("Meine Zahl ist kleiner")

		}
	}
	fmt.Println("Zu viele falsche Zahlen :-(")
}

func NumberGood(guess, expected int) bool {
	return guess == expected
}

func main() {
	GuessingGame()
}
