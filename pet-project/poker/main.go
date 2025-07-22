package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Welcome to Simple Poker! ===")
	fmt.Println("You start with 1000 chips. Good luck!")
	fmt.Println("WARNING: Folding means you lose any chips you've already bet (including blinds)!")
	
	game := newGame()
	scanner := bufio.NewScanner(os.Stdin)
	
	for !game.isGameOver() {
		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Printf("Starting new hand... (Dealer: %s)\n", game.players[game.dealer].name)
		
		// Post blinds first
		game.postBlinds()
		
		// Deal new hands
		game.dealHands()
		game.showPlayerHand()
		
		// Player's turn
		action := game.playerAction()
		
		switch action {
		case "1": // Bet/Raise
			game.playerBet()
		case "2": // Call
			callAmount := game.players[1].bet - game.players[0].bet
			if callAmount > game.players[0].chips {
				callAmount = game.players[0].chips
			}
			game.players[0].chips -= callAmount
			game.players[0].bet += callAmount
			game.pot += callAmount
			fmt.Printf("You call with %d additional chips. Total bet: %d, Pot: %d\n", 
				callAmount, game.players[0].bet, game.pot)
		case "3": // Fold
			fmt.Printf("You fold! You lose %d chips already bet.\n", game.players[0].bet)
			game.players[0].folded = true
		default:
			fmt.Printf("Invalid choice, you fold! You lose %d chips already bet.\n", game.players[0].bet)
			game.players[0].folded = true
		}
		
		// Computer's turn (if player didn't fold)
		if !game.players[0].folded {
			game.computerAction()
		}
		
		// Determine winner
		if game.players[0].folded {
			fmt.Printf("You folded. Computer wins the pot of %d chips!\n", game.pot)
			fmt.Printf("You lost %d chips from your bets/blinds.\n", game.players[0].bet)
			game.players[1].chips += game.pot
		} else if game.players[1].folded {
			fmt.Printf("Computer folded. You win the pot of %d chips!\n", game.pot)
			fmt.Printf("Computer lost %d chips from their bets/blinds.\n", game.players[1].bet)
			game.players[0].chips += game.pot
		} else {
			game.showdown()
		}
		
		// Show chip counts after hand
		fmt.Printf("\nChip counts after hand - You: %d, Computer: %d\n", 
			game.players[0].chips, game.players[1].chips)
		
		// Ask if player wants to continue
		if !game.isGameOver() {
			fmt.Print("\nPress Enter to continue to next hand (or type 'quit' to exit): ")
			scanner.Scan()
			input := scanner.Text()
			if input == "quit" {
				break
			}
		}
		
		// Reset for next round
		game.resetRound()
	}
	
	// Game over
	fmt.Println("\n=== GAME OVER ===")
	if game.players[0].chips > game.players[1].chips {
		fmt.Println("Congratulations! You won overall!")
	} else if game.players[0].chips < game.players[1].chips {
		fmt.Println("Computer wins overall! Better luck next time!")
	} else {
		fmt.Println("It's a tie!")
	}
	
	fmt.Printf("Final scores - You: %d, Computer: %d\n", game.players[0].chips, game.players[1].chips)
}
