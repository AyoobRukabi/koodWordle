package main

import (
	"fmt"
	"os"
	"strconv"

	"koodWordle/game"
	"koodWordle/io"
)

func main() {
	// handle missing command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number as command line argument")
		return
	}

	// parse argument
	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid command-line argument. Please launch with a valid number.")
		return
	}

	// load word list
	words, err := io.LoadWordList("wordle-words.txt")
	if err != nil {
		fmt.Println("Word list not found or error reading file.")
		return
	}
	
	// ask for username
	username := io.GetUsername()

	// check if index is valid
	if index < 0 || index >= len(words) {
		fmt.Println("Invalid word number.")
		fmt.Println("Press Enter to exit...")
		fmt.Scanln()
		return
	}

	secretWord := words[index]

	

	// start game
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
