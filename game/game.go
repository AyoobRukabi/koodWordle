// package game

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// const (
// 	Green  = "\u001B[32m"
// 	Yellow = "\u001B[33m"
// 	White  = "\u001B[37m"
// 	Reset  = "\u001B[0m"
// )

// // Play runs the game loop
// func Play(secretWord string) (bool, int) {
// 	attempts := 0
// 	maxAttempts := 6
// 	remainingLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")

// 	for attempts < maxAttempts {
// 		fmt.Print("Enter your guess: ")
// 		if !scanner.Scan() {
// 			break
// 		}
// 		guess := scanner.Text()

// 		// validate guess length
// 		if len(guess) != 5 {
// 			fmt.Println("Guess must be 5 letters.")
// 			continue
// 		}

// 		// validate lowercase letters
// 		if !isLowercase(guess) {
// 			fmt.Println("Your guess must only contain lowercase letters.")
// 			continue
// 		}

// 		attempts++
// 		feedback := GetFeedback(secretWord, guess)
// 		fmt.Printf("Feedback: %s\n", feedback)

// 		remainingLetters = updateRemainingLetters(remainingLetters, guess, secretWord)
// 		fmt.Printf("Remaining letters: %s\n", formatLetters(remainingLetters))
// 		fmt.Printf("Attempts remaining: %d\n", maxAttempts-attempts)

// 		if guess == secretWord {
// 			fmt.Printf("Congratulations! You guessed the word!\n")
// 			return true, attempts
// 		}
// 	}

// 	fmt.Printf("Sorry, you lost! The word was: %s\n", strings.ToUpper(secretWord))
// 	return false, attempts
// }

// // GetFeedback returns the colored feedback string
// func GetFeedback(secret, guess string) string {
// 	secret = strings.ToUpper(secret)
// 	feedback := ""
// 	for i := 0; i < 5; i++ {
// 		ch := strings.ToUpper(string(guess[i]))
// 		if guess[i] == strings.ToLower(string(secret[i]))[0] {
// 			feedback += Green + ch + Reset
// 		} else if strings.ContainsRune(secret, rune(ch[0])) {
// 			feedback += Yellow + ch + Reset
// 		} else {
// 			feedback += White + ch + Reset
// 		}
// 	}
// 	return feedback
// }

// // Update remaining letters after each guess
// func updateRemainingLetters(remaining, guess, secret string) string {
// 	result := ""
// 	guess = strings.ToUpper(guess)
// 	for _, r := range remaining {
// 		if !strings.ContainsRune(guess, r) || strings.ContainsRune(secret, r) {
// 			result += string(r)
// 		}
// 	}
// 	return result
// }

// // Format letters with spaces
// func formatLetters(s string) string {
// 	letters := []string{}
// 	for _, r := range s {
// 		letters = append(letters, string(r))
// 	}
// 	return strings.Join(letters, " ")
// }

// // Check if string contains only lowercase letters
// func isLowercase(s string) bool {
// 	for _, r := range s {
// 		if r < 'a' || r > 'z' {
// 			return false
// 		}
// 	}
// 	return true
// }




package game

import (
	"bufio"
	"fmt"
	"strings"
)

const maxAttempts = 6

// ANSI colors
const (
	green  = "\u001B[32m"
	yellow = "\u001B[33m"
	white  = "\u001B[37m"
	reset  = "\u001B[0m"
)

// Play starts the Wordle game.
func Play(scanner *bufio.Scanner, secret string, wordList []string) (bool, int) {
	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")

	attempts := 0
	secret = strings.ToUpper(secret)

	remainingLetters := make(map[rune]bool)
	for r := 'A'; r <= 'Z'; r++ {
		remainingLetters[r] = true
	}

	for attempts < maxAttempts {
		fmt.Print("Enter your guess:  ")
		if !scanner.Scan() {
			break
		}
		guess := scanner.Text()

		if len(guess) != 5 {
			fmt.Println("Your guess must be exactly 5 letters long.")
			continue
		}
		if !isLowercase(guess) {
			fmt.Println("Your guess must only contain lowercase letters.")
			continue
		}
		if !isWordInList(guess, wordList) {
			fmt.Println("Word not in list. Please enter a valid word.")
			continue
		}

		attempts++

		if strings.ToUpper(guess) == secret {
			fmt.Println("Congratulations! You've guessed the word correctly.")
			fmt.Print("Do you want to see your stats? (yes/no): ")
			return true, attempts
		}

		feedback := generateFeedback(secret, guess)
		fmt.Println("Feedback:", feedback)

		// Update remaining letters
		for _, ch := range strings.ToUpper(guess) {
			if !strings.ContainsRune(secret, ch) {
				remainingLetters[ch] = false
			}
		}

		// Print remaining letters
		fmt.Print("Remaining letters: ")
		for r := 'A'; r <= 'Z'; r++ {
			if remainingLetters[r] {
				fmt.Printf("%c ", r)
			}
		}
		fmt.Println()

		// Print remaining attempts
		fmt.Println("Attempts remaining: ", maxAttempts-attempts)

	}

	// Only after the loop ends
	if attempts >= maxAttempts {
		fmt.Printf("Game over. The correct word was: %s\n", strings.ToLower(secret))
		fmt.Print("Do you want to see your stats? (yes/no): ")
	}
	return false, attempts
}

func generateFeedback(secret, guess string) string {
	secret = strings.ToUpper(secret)
	guess = strings.ToUpper(guess)
	result := ""

	for i := 0; i < 5; i++ {
		if guess[i] == secret[i] {
			result += green + string(guess[i]) + reset
		} else if strings.ContainsRune(secret, rune(guess[i])) {
			result += yellow + string(guess[i]) + reset
		} else {
			result += white + string(guess[i]) + reset
		}
	}
	return result
}

func isLowercase(s string) bool {
	for _, c := range s {
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return true
}

func isWordInList(word string, list []string) bool {
	for _, w := range list {
		if w == word {
			return true
		}
	}
	return false
}
