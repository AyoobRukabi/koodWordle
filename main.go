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
		fmt.Println("Please provide a word index, e.g., go run . 10")
		return
	}

	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid index. Please provide a number.")
		return
	}

	// load word list
	words, err := io.LoadWordList("wordle-words.txt")
	if err != nil {
		fmt.Println("Word list not found or error reading file.")
		return
	}

	if index < 0 || index >= len(words) {
		fmt.Println("Index out of range.")
		return
	}

	secretWord := words[index]

	// ask for username
	username := io.GetUsername()
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
