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

	fmt.Print("Enter your guess: ")

	for attempts < maxAttempts {
		if !scanner.Scan() {
			break // handle EOF
		}
		guess := strings.TrimSpace(scanner.Text())

		// validate length
		if len(guess) != 5 {
			fmt.Println("Guess must be 5 letters.")
			fmt.Print("Enter your guess: ")
			continue
		}

		// validate lowercase letters only
		if !isLowercase(guess) {
			fmt.Println("Your guess must only contain lowercase letters.")
			fmt.Print("Enter your guess: ")
			continue
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
			fmt.Println("Congratulations! You guessed the word!")
			return true, attempts
		}

		// Prompt for next guess
		if attempts < maxAttempts {
			fmt.Print("Enter your guess: ")
		}
	}

	// After max attempts
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

// Update remaining letters
func updateRemainingLetters(remaining, guess, secret string) string {
	result := ""
	guess = strings.ToUpper(guess)
	for _, r := range remaining {
		if !strings.ContainsRune(guess, r) || strings.ContainsRune(secret, r) {
			result += string(r)
		}
	}
	return result
}

// Format remaining letters for display
func formatLetters(s string) string {
	letters := []string{}
	for _, r := range s {
		letters = append(letters, string(r))
	}
	return strings.Join(letters, " ")
}

// Check if string contains only lowercase letters
func isLowercase(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
