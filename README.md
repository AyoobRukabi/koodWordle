# koodWordle

A simple **Wordle-like game** written in Go. The program selects a word from a provided word list (`wordle-words.txt`) based on a line number you choose, and lets the player guess it.

## Description

koodWordle is a CLI-based word guessing game inspired by Wordle. Players provide a line number as an argument, and the program selects the word from that line in `wordle-words.txt` for the player to guess.

## How to Run

Make sure you have [Go](https://golang.org/dl/) installed.

1. Clone the repository:

```bash
git clone https://github.com/AyoobRukabi/koodWordle.git
cd koodWordle
Run the program with a line number as an argument:
go run main.go <line_number>
Example:
go run main.go 5
This will select the word on line 5 of wordle-words.txt as the word to guess.
Word List
The wordle-words.txt file should contain one word per line. Use the line number as the argument to choose a specific word from the list.

## Features
CLI-based Wordle game
Selects words based on line number from a file
Supports guessing via command-line arguments