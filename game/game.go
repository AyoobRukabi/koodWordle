package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	Green  = "\u001B[32m"
	Yellow = "\u001B[33m"
	White  = "\u001B[37m"
	Reset  = "\u001B[0m"
)

func Play(username, secretWord string) (bool, int) {
	attempts := 0
	maxAttempts := 6
	remainingLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	scanner := bufio.NewScanner(os.Stdin)
	secretWord = strings.ToLower(secretWord)

	for attempts < maxAttempts {
		fmt.Print("Enter your guess: ")
		if !scanner.Scan() {
			break
		}
		guess := scanner.Text()

		// Check 5 letters
		if len(guess) != 5 {
			fmt.Println("Your guess must be 5 letters.")
			continue
		}

		// Check lowercase
		if !isLowercase(guess) {
			fmt.Println("Your guess must only contain lowercase letters.")
			continue
		}

		attempts++

		// Generate feedback
		feedback := getFeedback(secretWord, guess)
		fmt.Printf("Feedback: %s\n", feedback)

		// Update remaining letters (remove only letters not in secret word)
		remainingLetters = updateRemainingLetters(remainingLetters, guess, secretWord)
		fmt.Printf("Remaining letters: %s\n", formatLetters(remainingLetters))
		fmt.Printf("Attempts remaining: %d\n", maxAttempts-attempts)

		if guess == secretWord {
			fmt.Println("Congratulations! You guessed the word!")
			return true, attempts
		}
	}

	fmt.Printf("Game over. The correct word was: %s\n", strings.ToUpper(secretWord))
	return false, attempts
}

func isLowercase(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func getFeedback(secret, guess string) string {
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
		if !strings.ContainsRune(guess, r) || strings.ContainsRune(secret, r) {
			result += string(r)
		}
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
