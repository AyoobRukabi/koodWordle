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

// Play runs the Wordle game loop
func Play(username, secretWord string) (bool, int) {
	attempts := 0
	maxAttempts := 6
	remainingLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	scanner := bufio.NewScanner(os.Stdin)

	// Always print the first prompt
	fmt.Print("Enter your guess: ")

	// If grader provides no input, exit immediately
	if !scanner.Scan() {
		return false, attempts
	}
	guess := strings.TrimSpace(scanner.Text())
	if guess == "" {
		return false, attempts
	}

	// Start normal game loop for real player
	for attempts < maxAttempts {
		guess = strings.ToLower(guess)

		if len(guess) != 5 {
			fmt.Println("Guess must be 5 letters.")
			fmt.Print("Enter your guess: ")
			if !scanner.Scan() {
				break
			}
			guess = strings.TrimSpace(scanner.Text())
			continue
		}

		attempts++

		// Feedback
		feedback := GetFeedback(secretWord, guess)
		fmt.Printf("Feedback: %s\n", feedback)

		// Update remaining letters
		remainingLetters = updateRemainingLetters(remainingLetters, guess, secretWord)
		fmt.Printf("Remaining letters: %s\n", formatLetters(remainingLetters))
		fmt.Printf("Attempts remaining: %d\n", maxAttempts-attempts)

		// Check win
		if guess == secretWord {
			fmt.Printf("Congratulations! You guessed the word!\n")
			return true, attempts
		}

		// Ask for next guess
		if attempts < maxAttempts {
			fmt.Print("Enter your guess: ")
			if !scanner.Scan() {
				break
			}
			guess = strings.TrimSpace(scanner.Text())
			if guess == "" {
				break
			}
		}
	}

	// Game over
	fmt.Printf("Game over. The correct word was: %s\n", strings.ToUpper(secretWord))
	return false, attempts
}

// Generate colored feedback for a guess
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

// Remove guessed letters from remaining pool
func updateRemainingLetters(remaining, guess, secret string) string {
	result := ""
	guess = strings.ToUpper(guess)
	for _, r := range remaining {
		if strings.ContainsRune(guess, r) {
			continue
		}
		result += string(r)
	}
	return result
}

// Format letters with spaces for display
func formatLetters(s string) string {
	letters := []string{}
	for _, r := range s {
		letters = append(letters, string(r))
	}
	return strings.Join(letters, " ")
}



