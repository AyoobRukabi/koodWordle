package main

import (
	"fmt"
	"os"
	"strconv"

	"koodWordle/game"
	"koodWordle/io"
)

func main() {
	// handle command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number as command line argument")
		return
	}

	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid command-line argument. Please launch with a valid number.")
		return
	}

	// ask for username first (so output always starts with this prompt)
	username := io.GetUsername()

	// load word list
	words, err := io.LoadWordList("wordle-words.txt")
	if err != nil {
		fmt.Println("Word list not found or error reading file.")
		return
	}

	// check validity of the word number
	if index < 0 || index >= len(words) {
		fmt.Println("Invalid word number.")
		fmt.Println("Press Enter to exit...")
		fmt.Scanln() // wait for Enter
		return
	}

	// valid secret word
	secretWord := words[index]

	fmt.Printf("Welcome to Wordle! Guess the 5-letter word.\n")

	// start the game
	win, attempts := game.Play(username, secretWord)

	// save stats
	result := "loss"
	if win {
		result = "win"
	}
	io.SaveStats(username, secretWord, attempts, result)

	// show stats
	io.ShowStats(username)
}
