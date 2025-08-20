package main

import (
	"fmt"
	"os"
	"strconv"

	"koodWordle/game"
	"koodWordle/io"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number as command line argument")
		return
	}

	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid command-line argument. Please launch with a valid number.")
		return
	}

	words, err := io.LoadWordList("wordle-words.txt")
	if err != nil {
		fmt.Println("Word list not found or error reading file.")
		return
	}

	if index < 0 || index >= len(words) {
		// Prompt username even if index invalid
		fmt.Print("Enter your username: ")
		fmt.Println("Invalid word number.")
		fmt.Println("Press Enter to exit...")
		fmt.Scanln()
		return
	}

	secretWord := words[index]

	username := io.GetUsername()
	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")

	win, attempts := game.Play(username, secretWord)

	result := "loss"
	if win {
		result = "win"
	}
	io.SaveStats(username, secretWord, attempts, result)

	io.ShowStats(username)
}
