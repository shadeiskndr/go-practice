package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

// Card represents a single playing card
type card struct {
	value int    // 2-14 (where 11=Jack, 12=Queen, 13=King, 14=Ace)
	suit  string
}

// HandRank represents the strength of a poker hand
type handRank struct {
	rank     int    // 1=High Card, 2=Pair, 3=Two Pair, etc.
	rankName string
	values   []int  // For tie-breaking
}

// Player represents a poker player
type player struct {
	name   string
	hand   deck
	chips  int
	bet    int
	folded bool
}

// Game represents the poker game state
type game struct {
	players    []player
	deck       deck
	pot        int
	round      string
	smallBlind int
	bigBlind   int
	dealer     int // 0 or 1, alternates each hand
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert string card to card struct
func parseCard(cardStr string) card {
	parts := strings.Split(cardStr, " of ")
	valueName := parts[0]
	suit := parts[1]
	
	var value int
	switch valueName {
	case "Two":
		value = 2
	case "Three":
		value = 3
	case "Four":
		value = 4
	case "Five":
		value = 5
	case "Six":
		value = 6
	case "Seven":
		value = 7
	case "Eight":
		value = 8
	case "Nine":
		value = 9
	case "Ten":
		value = 10
	case "Jack":
		value = 11
	case "Queen":
		value = 12
	case "King":
		value = 13
	case "Ace":
		value = 14
	}
	
	return card{value: value, suit: suit}
}

// Convert deck to cards for evaluation
func (d deck) toCards() []card {
	cards := make([]card, len(d))
	for i, cardStr := range d {
		cards[i] = parseCard(cardStr)
	}
	return cards
}

// Evaluate poker hand strength
func evaluateHand(cards []card) handRank {
	// Sort cards by value for easier evaluation
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].value > cards[j].value
	})
	
	// Count values and suits
	valueCounts := make(map[int]int)
	suitCounts := make(map[string]int)
	values := make([]int, len(cards))
	
	for i, card := range cards {
		valueCounts[card.value]++
		suitCounts[card.suit]++
		values[i] = card.value
	}
	
	// Check for flush
	isFlush := false
	for _, count := range suitCounts {
		if count == 5 {
			isFlush = true
			break
		}
	}
	
	// Check for straight
	isStraight := false
	if len(valueCounts) == 5 {
		// Check normal straight
		if values[0]-values[4] == 4 {
			isStraight = true
		}
		// Check A-2-3-4-5 straight (wheel)
		if values[0] == 14 && values[1] == 5 && values[2] == 4 && values[3] == 3 && values[4] == 2 {
			isStraight = true
			values[0] = 1 // Ace low in wheel
			sort.Ints(values)
		}
	}
	
	// Count pairs, trips, etc.
	var pairs []int
	var trips []int
	var quads []int
	
	for value, count := range valueCounts {
		switch count {
		case 2:
			pairs = append(pairs, value)
		case 3:
			trips = append(trips, value)
		case 4:
			quads = append(quads, value)
		}
	}
	
	// Sort for consistent ordering
	sort.Sort(sort.Reverse(sort.IntSlice(pairs)))
	sort.Sort(sort.Reverse(sort.IntSlice(trips)))
	
	// Determine hand rank
	if isStraight && isFlush {
		if values[0] == 14 && values[1] == 13 { // Royal flush
			return handRank{rank: 10, rankName: "Royal Flush", values: values}
		}
		return handRank{rank: 9, rankName: "Straight Flush", values: values}
	}
	
	if len(quads) > 0 {
		return handRank{rank: 8, rankName: "Four of a Kind", values: append(quads, getKickers(values, append(quads, trips...))...)}
	}
	
	if len(trips) > 0 && len(pairs) > 0 {
		return handRank{rank: 7, rankName: "Full House", values: append(trips, pairs...)}
	}
	
	if isFlush {
		return handRank{rank: 6, rankName: "Flush", values: values}
	}
	
	if isStraight {
		return handRank{rank: 5, rankName: "Straight", values: values}
	}
	
	if len(trips) > 0 {
		return handRank{rank: 4, rankName: "Three of a Kind", values: append(trips, getKickers(values, trips)...)}
	}
	
	if len(pairs) >= 2 {
		return handRank{rank: 3, rankName: "Two Pair", values: append(pairs, getKickers(values, pairs)...)}
	}
	
	if len(pairs) == 1 {
		return handRank{rank: 2, rankName: "One Pair", values: append(pairs, getKickers(values, pairs)...)}
	}
	
	return handRank{rank: 1, rankName: "High Card", values: values}
}

// Get kicker cards (cards not part of the main hand)
func getKickers(allValues []int, usedValues []int) []int {
	used := make(map[int]bool)
	for _, v := range usedValues {
		used[v] = true
	}
	
	var kickers []int
	for _, v := range allValues {
		if !used[v] {
			kickers = append(kickers, v)
		}
	}
	
	sort.Sort(sort.Reverse(sort.IntSlice(kickers)))
	return kickers
}

// Compare two hands - returns 1 if hand1 wins, -1 if hand2 wins, 0 if tie
func compareHands(hand1, hand2 handRank) int {
	if hand1.rank > hand2.rank {
		return 1
	}
	if hand1.rank < hand2.rank {
		return -1
	}
	
	// Same rank, compare values
	for i := 0; i < len(hand1.values) && i < len(hand2.values); i++ {
		if hand1.values[i] > hand2.values[i] {
			return 1
		}
		if hand1.values[i] < hand2.values[i] {
			return -1
		}
	}
	
	return 0 // Tie
}

func newGame() *game {
	g := &game{
		players: []player{
			{name: "You", chips: 1000, folded: false},
			{name: "Computer", chips: 1000, folded: false},
		},
		deck:       newDeck(),
		pot:        0,
		round:      "pre-flop",
		smallBlind: 25,
		bigBlind:   50,
		dealer:     0, // Player starts as dealer
	}
	g.deck.shuffle()
	return g
}

func (g *game) postBlinds() {
	fmt.Println("\n=== Posting Blinds ===")
	
	// Determine who posts what based on dealer position
	var smallBlindPlayer, bigBlindPlayer int
	if g.dealer == 0 {
		// Player is dealer, so player posts small blind, computer posts big blind
		smallBlindPlayer = 0
		bigBlindPlayer = 1
	} else {
		// Computer is dealer, so computer posts small blind, player posts big blind
		smallBlindPlayer = 1
		bigBlindPlayer = 0
	}
	
	// Post small blind
	smallBlindAmount := g.smallBlind
	if smallBlindAmount > g.players[smallBlindPlayer].chips {
		smallBlindAmount = g.players[smallBlindPlayer].chips
	}
	g.players[smallBlindPlayer].chips -= smallBlindAmount
	g.players[smallBlindPlayer].bet = smallBlindAmount
	g.pot += smallBlindAmount
	fmt.Printf("%s posts small blind: %d chips\n", g.players[smallBlindPlayer].name, smallBlindAmount)
	
	// Post big blind
	bigBlindAmount := g.bigBlind
	if bigBlindAmount > g.players[bigBlindPlayer].chips {
		bigBlindAmount = g.players[bigBlindPlayer].chips
	}
	g.players[bigBlindPlayer].chips -= bigBlindAmount
	g.players[bigBlindPlayer].bet = bigBlindAmount
	g.pot += bigBlindAmount
	fmt.Printf("%s posts big blind: %d chips\n", g.players[bigBlindPlayer].name, bigBlindAmount)
	
	fmt.Printf("Pot after blinds: %d chips\n", g.pot)
}

func (g *game) dealHands() {
	for i := range g.players {
		hand, remaining := deal(g.deck, 5)
		g.players[i].hand = hand
		g.deck = remaining
	}
}

func (g *game) showPlayerHand() {
	fmt.Println("\n=== Your Hand ===")
	fmt.Println(g.players[0].hand.toString())
	
	// Show hand strength
	playerCards := g.players[0].hand.toCards()
	playerRank := evaluateHand(playerCards)
	fmt.Printf("Your hand: %s\n", playerRank.rankName)
	
	fmt.Printf("Your chips: %d\n", g.players[0].chips)
	fmt.Printf("Current pot: %d\n", g.pot)
	fmt.Printf("Your current bet: %d\n", g.players[0].bet)
}

func (g *game) playerAction() string {
	fmt.Println("\nWhat would you like to do?")
	fmt.Println("1. Bet/Raise")
	fmt.Println("2. Call")
	fmt.Println("3. Fold (WARNING: You'll lose your blinds/bets!)")
	fmt.Print("Enter your choice (1-3): ")
	
	var choice string
	fmt.Scanln(&choice)
	return choice
}

func (g *game) playerBet() {
	currentBet := g.players[1].bet // Computer's current bet
	minRaise := currentBet + g.bigBlind
	
	fmt.Printf("Current bet to call: %d\n", currentBet)
	fmt.Printf("Minimum raise: %d\n", minRaise)
	fmt.Printf("How much would you like to bet? (Max: %d): ", g.players[0].chips+g.players[0].bet)
	
	var betAmount int
	fmt.Scanln(&betAmount)
	
	// Calculate additional chips needed
	additionalBet := betAmount - g.players[0].bet
	
	if additionalBet > g.players[0].chips {
		additionalBet = g.players[0].chips
		betAmount = g.players[0].bet + additionalBet
		fmt.Printf("Betting all remaining chips. Total bet: %d\n", betAmount)
	}
	
	g.players[0].chips -= additionalBet
	g.players[0].bet = betAmount
	g.pot += additionalBet
	fmt.Printf("You bet %d chips total. Pot is now %d\n", betAmount, g.pot)
}

func (g *game) computerAction() {
	// Simple AI based on hand strength
	computerCards := g.players[1].hand.toCards()
	computerRank := evaluateHand(computerCards)
	
	currentBet := g.players[0].bet
	callAmount := currentBet - g.players[1].bet
	
	// AI decision based on hand strength
	var action int
	if computerRank.rank >= 6 { // Flush or better - always bet/raise
		action = 2
	} else if computerRank.rank >= 3 { // Two pair or better - usually bet/call
		action = rand.Intn(2) + 1 // 1 or 2
	} else if computerRank.rank == 2 { // One pair - mixed strategy
		action = rand.Intn(3)
	} else { // High card - usually fold
		if rand.Intn(4) == 0 { // 25% chance to bluff
			action = rand.Intn(2) + 1
		} else {
			action = 0 // fold
		}
	}
	
	switch action {
	case 0: // Fold
		fmt.Printf("Computer folds! (Loses %d chips already bet)\n", g.players[1].bet)
		g.players[1].folded = true
	case 1: // Call
		if callAmount > g.players[1].chips {
			callAmount = g.players[1].chips
		}
		g.players[1].chips -= callAmount
		g.players[1].bet += callAmount
		g.pot += callAmount
		fmt.Printf("Computer calls with %d chips. Pot is now %d\n", callAmount, g.pot)
	case 2: // Bet/Raise
		raiseAmount := 50 + (computerRank.rank * 20) // Bet more with better hands
		totalBet := currentBet + raiseAmount
		additionalBet := totalBet - g.players[1].bet
		
		if additionalBet > g.players[1].chips {
			additionalBet = g.players[1].chips
			totalBet = g.players[1].bet + additionalBet
		}
		
		g.players[1].chips -= additionalBet
		g.players[1].bet = totalBet
		g.pot += additionalBet
		fmt.Printf("Computer raises to %d chips. Pot is now %d\n", totalBet, g.pot)
	}
}

func (g *game) showdown() {
	fmt.Println("\n=== SHOWDOWN ===")
	
	playerCards := g.players[0].hand.toCards()
	computerCards := g.players[1].hand.toCards()
	
	playerRank := evaluateHand(playerCards)
	computerRank := evaluateHand(computerCards)
	
	fmt.Printf("Your hand: %s (%s)\n", g.players[0].hand.toString(), playerRank.rankName)
	fmt.Printf("Computer hand: %s (%s)\n", g.players[1].hand.toString(), computerRank.rankName)
	
	result := compareHands(playerRank, computerRank)
	
	if result > 0 {
		fmt.Println("You win the hand!")
		g.players[0].chips += g.pot
	} else if result < 0 {
		fmt.Println("Computer wins the hand!")
		g.players[1].chips += g.pot
	} else {
		fmt.Println("It's a tie! Pot is split.")
		splitPot := g.pot / 2
		g.players[0].chips += splitPot
		g.players[1].chips += (g.pot - splitPot) // Handle odd pots
	}
	
	fmt.Printf("Your chips: %d\n", g.players[0].chips)
	fmt.Printf("Computer chips: %d\n", g.players[1].chips)
}

func (g *game) resetRound() {
	g.pot = 0
	g.round = "pre-flop"
	
	// Switch dealer
	g.dealer = 1 - g.dealer
	
	for i := range g.players {
		g.players[i].bet = 0
		g.players[i].folded = false
		g.players[i].hand = deck{}
	}
	g.deck = newDeck()
	g.deck.shuffle()
}

func (g *game) isGameOver() bool {
	return g.players[0].chips <= 0 || g.players[1].chips <= 0
}
