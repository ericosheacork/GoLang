package main

func main() {
	cards := readDeckFromFile("deck_box")
	// // cards = append(cards, "Six of Spades")
	// // // or card:= "Ace of Spades" only use the syntax when first declaring the var
	// // fmt.Println(cards)

	// // cards.print()

	// //returns 2 values so we use syntax to capture it

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// greeting := "Hi There"
	// fmt.Println([]byte(greeting))
	// fm.Println(cards.toString())t
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
