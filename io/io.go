// package io

// import (
// 	"bufio"
// 	"encoding/csv"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// // Load words from file
// func LoadWordList(filename string) ([]string, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var words []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		words = append(words, scanner.Text())
// 	}
// 	return words, scanner.Err()
// }

// // Ask for username
// func GetUsername() string {
// 	fmt.Print("Enter your username: ")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	username := "Anonymous"
// 	if scanner.Scan() {
// 		username = strings.TrimSpace(scanner.Text())
// 	}
// 	return username
// }

// // Save game stats
// func SaveStats(username, secretWord string, attempts int, result string) error {
// 	file, err := os.OpenFile("stats.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	record := []string{
// 		username,
// 		secretWord,
// 		fmt.Sprintf("%d", attempts),
// 		result,
// 	}

// 	return writer.Write(record)
// }

// // Load stats
// func LoadStats(username string) (gamesPlayed, gamesWon, totalAttempts int, err error) {
// 	file, err := os.Open("stats.csv")
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return 0, 0, 0, nil
// 		}
// 		return 0, 0, 0, err
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		return 0, 0, 0, err
// 	}

// 	for _, record := range records {
// 		if len(record) < 4 {
// 			continue
// 		}
// 		if record[0] == username {
// 			gamesPlayed++
// 			attempts := 0
// 			fmt.Sscanf(record[2], "%d", &attempts)
// 			totalAttempts += attempts
// 			if record[3] == "win" {
// 				gamesWon++
// 			}
// 		}
// 	}

// 	return gamesPlayed, gamesWon, totalAttempts, nil
// }

// // Show stats
// func ShowStats(username string) {
// 	fmt.Print("Do you want to see your stats? (yes/no): ")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	if !scanner.Scan() {
// 		return
// 	}
// 	answer := strings.TrimSpace(scanner.Text())
// 	if strings.ToLower(answer) != "yes" {
// 		return
// 	}

// 	gamesPlayed, gamesWon, totalAttempts, _ := LoadStats(username)

// 	avgAttempts := 0.0
// 	if gamesPlayed > 0 {
// 		avgAttempts = float64(totalAttempts) / float64(gamesPlayed)
// 	}

// 	fmt.Printf("Stats for %s:\n", username)
// 	fmt.Printf("Games played: %d\n", gamesPlayed)
// 	fmt.Printf("Games won: %d\n", gamesWon)
// 	fmt.Printf("Average attempts per game: %.2f\n", avgAttempts)
// 	fmt.Print("Press Enter to exit...")
// 	scanner.Scan() // wait for Enter
// }


package io

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// LoadWords reads a word list from a file.
func LoadWords(filename string) ([]string, error) {
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

// SaveStats appends a game record to CSV.
func SaveStats(filename, username, secret string, attempts int, status string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error saving stats:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{username, secret, strconv.Itoa(attempts), status}
	writer.Write(record)
}

type Stats struct {
	User          string
	GamesPlayed   int
	GamesWon      int
	TotalAttempts int
}

// LoadStats loads stats for a given user.
func LoadStats(filename, username string) (Stats, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Stats{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return Stats{}, err
	}

	stats := Stats{User: username}
	for _, r := range records {
		if r[0] == username {
			stats.GamesPlayed++
			if r[3] == "win" {
				stats.GamesWon++
			}
			att, _ := strconv.Atoi(r[2])
			stats.TotalAttempts += att
		}
	}
	return stats, nil
}

// Print prints the user stats to stdout.
func (s Stats) Print() {
	avg := 0.0
	if s.GamesPlayed > 0 {
		avg = float64(s.TotalAttempts) / float64(s.GamesPlayed)
	}
	fmt.Printf("Stats for %s:\n", s.User)
	fmt.Println("Games played:", s.GamesPlayed)
	fmt.Println("Games won:", s.GamesWon)
	fmt.Printf("Average attempts per game: %.2f\n", avg)
}
