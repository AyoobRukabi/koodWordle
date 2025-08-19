package io

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

// Load words from file
func LoadWordList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

// Ask for username
func GetUsername() string {
	fmt.Print("Enter your username: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return "Anonymous"
}

// Save game stats
func SaveStats(username, secretWord string, attempts int, result string) error {
	file, err := os.OpenFile("stats.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		username,
		secretWord,
		fmt.Sprintf("%d", attempts),
		result,
	}

	return writer.Write(record)
}

// Load stats
func LoadStats(username string) (gamesPlayed, gamesWon, totalAttempts int, err error) {
	file, err := os.Open("stats.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return 0, 0, 0, nil
		}
		return 0, 0, 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, 0, 0, err
	}

	for _, record := range records {
		if len(record) < 4 {
			continue
		}
		if record[0] == username {
			gamesPlayed++
			attempts := 0
			fmt.Sscanf(record[2], "%d", &attempts)
			totalAttempts += attempts
			if record[3] == "win" {
				gamesWon++
			}
		}
	}

	return gamesPlayed, gamesWon, totalAttempts, nil
}

// Show stats
func ShowStats(username string) {
	fmt.Print("Do you want to see your stats? (yes/no): ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	if scanner.Text() != "yes" {
		return
	}

	gamesPlayed, gamesWon, totalAttempts, _ := LoadStats(username)

	avgAttempts := 0.0
	if gamesPlayed > 0 {
		avgAttempts = float64(totalAttempts) / float64(gamesPlayed)
	}

	fmt.Printf("Stats for %s:\n", username)
	fmt.Printf("Games played: %d\n", gamesPlayed)
	fmt.Printf("Games won: %d\n", gamesWon)
	fmt.Printf("Average attempts per game: %.2f\n", avgAttempts)
	fmt.Println("Press Enter to exit...")
	scanner.Scan() // wait for Enter
}
