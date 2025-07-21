package main

import "fmt"

func main (){

	deck := newDeck()
	deck.shuffle()
	deck.saveToFile("all_cards.txt")
	newDeckFromFile("all_cards.txt")
	hand, remaining := deal(deck, 5)
	fmt.Println("=============-hand-==============")
	fmt.Println(hand.toString())
	hand.saveToFile("my_hand.txt")
	fmt.Println("===========-remaining-===========")
	fmt.Println(remaining.toString())
	remaining.saveToFile("remaining_cards.txt")

	// cards := deck{"Six of Spades,", newCard()}
	// cards = append(cards, "King of Clubs,")
	// fmt.Println(cards)

	// fmt.Println("==============================")

	// for i, cards := range cards {
	// 	fmt.Println(i, cards)
	// }

	// fmt.Println("==============================")

	// for i := range cards {
	// 	fmt.Println(cards[i])
	// }

	// fmt.Println("==============================")

	// cards.print()

	// fmt.Println("==============================")
	
	// for i := range newDeck() {
	// 	fmt.Println(newDeck()[i])
	// }

	// fmt.Println("==============================")

	// hand, remainingCards := deal(newDeck(), 5)
	// hand.print()

	// fmt.Println("------------------------------")
	
	// remainingCards.print()

	// fmt.Println("==============================")

	// fmt.Println(hand.toString())

	// fmt.Println("------------------------------")

	// fmt.Println(remainingCards.toString())

	// fmt.Println("==============================")
}

func newCard() string {
	return "Five of Diamonds,"
}