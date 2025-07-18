package main

import "fmt"

func main (){
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

	fmt.Println("==============================")

	hand, remainingCards := deal(newDeck(), 5)
	hand.print()

	fmt.Println("------------------------------")
	
	remainingCards.print()

	fmt.Println("==============================")
}

func newCard() string {
	return "Five of Diamonds,"
}