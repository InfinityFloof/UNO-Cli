package main

// Import used packages
import (
	"fmt"
	"math/rand"
	"time"
)

// Rules
var (
	O7        bool = false
	JumpIn    bool = false
	Challenge bool = false
)

// This is a struct that shows the schematics of what a card should be built on.
type card struct {
	color       string
	name        string
	amountCards int
}

type handCard struct {
	color       string
	name        string
}

// This is a struct that individual player data
type playerChar struct {
	name         string
	amountInHand int
	cardsInHand  [30]handCard
	isUNO        bool
}

// Create players
var player playerChar
var computer1 playerChar
var computer2 playerChar
var computer3 playerChar

// This is our deck, here's where all the cards in the UNO game is stored.
var deckSchematic = []card{

	/*
		Cards in a UNO deck

		To explain the structure

		{"Color", "Type", amount}
	*/

	// Black/Universal cards
	{"Black", "Wildcard", 4}, {"Black", "Plus4", 4},

	// Zero cards
	{"Red", "Zero", 1}, {"Yellow", "Zero", 1}, {"Green", "Zero", 1}, {"Blue", "Zero", 1},

	// +2 cards
	{"Red", "Plus2", 2}, {"Yellow", "Plus2", 2}, {"Green", "Plus2", 2}, {"Blue", "Plus2", 2},

	// Reverse cards
	{"Red", "Reverse", 2}, {"Yellow", "Reverse", 2}, {"Green", "Reverse", 2}, {"Blue", "Reverse", 2},

	// Block cards
	{"Red", "Block", 2}, {"Yellow", "Block", 2}, {"Green", "Block", 2}, {"Blue", "Block", 2},

	// Red cards
	{"Red", "One", 2}, {"Red", "Two", 2}, {"Red", "Three", 2}, {"Red", "Four", 2}, {"Red", "Five", 2}, {"Red", "Six", 2}, {"Red", "Seven", 2}, {"Red", "Eight", 2}, {"Red", "Nine", 2},

	// Yellow cards
	{"Yellow", "One", 2}, {"Yellow", "Two", 2}, {"Yellow", "Three", 2}, {"Yellow", "Four", 2}, {"Yellow", "Five", 2}, {"Yellow", "Six", 2}, {"Yellow", "Seven", 2}, {"Yellow", "Eight", 2}, {"Yellow", "Nine", 2},

	// Green cards
	{"Green", "One", 2}, {"Green", "One", 2}, {"Green", "One", 2}, {"Green", "One", 2}, {"Green", "Five", 2}, {"Green", "Six", 2}, {"Green", "Seven", 2}, {"Green", "Eight", 2}, {"Green", "Nine", 2},

	// Blue cards
	{"Red", "One", 2}, {"Red", "Two", 2}, {"Red", "Three", 2}, {"Red", "Four", 2}, {"Red", "Five", 2}, {"Red", "Six", 2}, {"Red", "Seven", 2}, {"Red", "Eight", 2}, {"Red", "Nine", 2},
}

// Copies the deckSchematic to deck so we can vary the contents during the progress of the game
var deck = deckSchematic

// Get total cards
func getCardTotal() int {
	a := 0
	for i := 0; i < len(deck); i++ {
		a = a + deck[i].amountCards
	}
	return a
}

// Generate an integer
func randomInt(a int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(a)
}

// Draw a card
func drawCard(a string) {
	for true {
		var b int = randomInt(len(deckSchematic)-1)

		// Checks that we dont have a full hand
		switch a {
		case "player":
			if player.amountInHand==30 {
				break
			}
		case "computer1":
			if computer1.amountInHand==30 {
				break
			}
		case "computer2":
			if computer2.amountInHand==30 {
				break
			}
		case "computer3":
			if computer3.amountInHand==30 {
				break
			}
		}

		// Checks that we have the card in the deck
		if deck[b].amountCards!=0 {

			// Selects player and gives them the card
			switch a {
			case "player":
				player.cardsInHand[player.amountInHand] = handCard{deck[b].color, deck[b].name}
				deck[b].amountCards--
				player.amountInHand++
			case "computer1":
				computer1.cardsInHand[computer1.amountInHand] = handCard{deck[b].color, deck[b].name}
				deck[b].amountCards--
				computer1.amountInHand++
			case "computer2":
				computer2.cardsInHand[computer2.amountInHand] = handCard{deck[b].color, deck[b].name}
				deck[b].amountCards--
				computer2.amountInHand++
			case "computer3":
				computer3.cardsInHand[computer3.amountInHand] = handCard{deck[b].color, deck[b].name}
				deck[b].amountCards--
				computer3.amountInHand++
			}
			break
		}
	}
}

func main() {
	// Hands everyone 7 cards
	for i := 0; i <= 7; i++{
		drawCard("player")
		drawCard("computer1")
		drawCard("computer2")
		drawCard("computer3")
	}	
	fmt.Println(player.cardsInHand)
}
