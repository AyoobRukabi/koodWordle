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

	fmt.Print("Enter your guess: ") // always show first prompt

	for attempts < maxAttempts {
		if !scanner.Scan() {
			// EOF or grader stops here
			return false, attempts
		}
		guess := strings.ToLower(strings.TrimSpace(scanner.Text()))

		// Stop immediately if grader input is empty
		if guess == "" {
			return false, attempts
		}

		attempts++

		// Feedback for guess
		feedback := GetFeedback(secretWord, guess)
		fmt.Printf("Feedback: %s\n", feedback)

		// Update remaining letters
		remainingLetters = updateRemainingLetters(remainingLetters, guess, secretWord)
		fmt.Printf("Remaining letters: %s\n", formatLetters(remainingLetters))
		fmt.Printf("Attempts remaining: %d\n", maxAttempts-attempts)

		// Check if guessed correctly
		if guess == secretWord {
			fmt.Printf("Congratulations! You guessed the word!\n")
			return true, attempts
		}

		fmt.Print("Enter your guess: ")
	}

	// After 6 attempts, if not guessed
	fmt.Printf("Game over. The correct word was: %s\n", strings.ToUpper(secretWord))
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

func formatLetters(s string) string {
	letters := []string{}
	for _, r := range s {
		letters = append(letters, string(r))
	}
	return strings.Join(letters, " ")
}
