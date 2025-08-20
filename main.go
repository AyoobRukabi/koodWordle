package main

import (
    "fmt"
    "os"
    "strconv"

    "koodWordle/game"
    "koodWordle/io"
)

func main() {
    // ask for username first (always)
    fmt.Print("Enter your username: ")
    username := io.GetUsername()

    // handle command-line arguments
    if len(os.Args) < 2 {
        fmt.Println("Invalid command-line argument. Please launch with a valid number.")
        return
    }

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

    if index < 0 || index >= len(words) {
        fmt.Println("Invalid word number.")
        fmt.Println("Press Enter to exit...")
        fmt.Scanln() // wait for Enter
        return
    }

    secretWord := words[index]

    // start the game
    fmt.Println("Welcome to Wordle! Guess the 5-letter word.")
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
