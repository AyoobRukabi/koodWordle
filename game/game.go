package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Green  = "\u001B[32m"
	Yellow = "\u001B[33m"
	White  = "\u001B[37m"
	Reset  = "\u001B[0m"
)

// Play runs the game loop
func Play(username, secretWord string) (bool, int) {
	attempts := 0
	maxAttempts := 6
	remainingLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	scanner := bufio.NewScanner(os.Stdin)

	for attempts < maxAttempts {
		fmt.Print("Enter your guess: ")
		if !scanner.Scan() {
			break // handle EOF (Ctrl+D)
		}
		guess := strings.ToLower(scanner.Text())

		if len(guess) != 5 {
			fmt.Println("Guess must be 5 letters.")
			continue
		}

		attempts++

		feedback := GetFeedback(secretWord, guess)
		fmt.Printf("Feedback: %s\n", feedback)

		remainingLetters = updateRemainingLetters(remainingLetters, guess, secretWord)
		fmt.Printf("Remaining letters: %s\n", formatLetters(remainingLetters))
		fmt.Printf("Attempts remaining: %d\n", maxAttempts-attempts)

		if guess == secretWord {
			fmt.Printf("Congratulations, %s! You guessed the word!\n", username)
			return true, attempts
		}
	}

	fmt.Printf("Sorry, you lost! The word was: %s\n", strings.ToUpper(secretWord))
	return false, attempts
}

// Generate feedback string
func GetFeedback(secret, guess string) string {
	feedback := ""
	for i := 0; i < 5; i++ {
		ch := strings.ToUpper(string(guess[i]))
		if guess[i] == secret[i] {
			feedback += Green + ch + Reset
		} else if strings.ContainsRune(secret, rune(guess[i])) {
			feedback += Yellow + ch + Reset
		} else {
			feedback += White + ch + Reset
		}
	}
	return feedback
}



func updateRemainingLetters(remaining, guess, secret string) string {
	result := ""
	guess = strings.ToUpper(guess)
	for _, r := range remaining {
		// remove any guessed letter from the remaining pool
		if strings.ContainsRune(guess, r) {
			continue
		}
		result += string(r)
	}
	return result
}













func formatLetters(s string) string {
	letters := []string{}
	for _, r := range s {
		letters = append(letters, string(r))
	}
	return strings.Join(letters, " ")
}
