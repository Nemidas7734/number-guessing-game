package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func RulesOfGame() {
	fmt.Println("\nWelcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have a limited number of chances to guess the correct number.")
	fmt.Println("\nPlease select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (2 chances)")
	fmt.Print("\nEnter your choice (1-3): ")
}

func playAgain() bool {
	var response string
	fmt.Print("\nDo you want to play again? (Y/N): ")
	fmt.Scan(&response)
	if response == "Y" || response == "y" {
		return true
	}
	return false
}

func EasyGame(scanner *bufio.Scanner, highScores map[string]int) {
	fmt.Println("Great! You have selected the Easy difficulty level.")
	fmt.Println("Let's start the game!")

	computerNo := rand.Intn(100) + 1

	for i := 1; i <= 10; i++ {
		fmt.Print("\nEnter your guess: ")
		scanner.Scan()
		guessNo, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid number, please try again.")
			i--
			continue
		}

		if guessNo == computerNo {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", i)
			if highScores["easy"] == 0 || i < highScores["easy"] {
				highScores["easy"] = i
				fmt.Printf("New high score for Easy difficulty: %d attempts!\n", highScores["easy"])
			}
			return
		} else if guessNo < computerNo {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guessNo)
		} else {
			fmt.Printf("Incorrect! The number is less than %d.\n", guessNo)
		}
	}
	fmt.Printf("\nSorry! You've run out of attempts. The correct number was %d.\n", computerNo)
}

func MediumGame(scanner *bufio.Scanner, highScores map[string]int) {
	fmt.Println("Great! You have selected the Medium difficulty level.")
	fmt.Println("Let's start the game!")

	computerNo := rand.Intn(100) + 1

	for i := 1; i <= 5; i++ {
		fmt.Print("\nEnter your guess: ")
		scanner.Scan()
		guessNo, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid number, please try again.")
			i--
			continue
		}

		if guessNo == computerNo {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", i)
			if highScores["medium"] == 0 || i < highScores["medium"] {
				highScores["medium"] = i
				fmt.Printf("New high score for Medium difficulty: %d attempts!\n", highScores["medium"])
			}
			return
		} else if guessNo < computerNo {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guessNo)
		} else {
			fmt.Printf("Incorrect! The number is less than %d.\n", guessNo)
		}
	}
	fmt.Printf("\nSorry! You've run out of attempts. The correct number was %d.\n", computerNo)
}

func HardGame(scanner *bufio.Scanner, highScores map[string]int) {
	fmt.Println("Great! You have selected the Hard difficulty level.")
	fmt.Println("Let's start the game!")

	computerNo := rand.Intn(100) + 1

	for i := 1; i <= 2; i++ {
		fmt.Print("\nEnter your guess: ")
		scanner.Scan()
		guessNo, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid number, please try again.")
			i--
			continue
		}

		if guessNo == computerNo {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", i)
			if highScores["hard"] == 0 || i < highScores["hard"] {
				highScores["hard"] = i
				fmt.Printf("New high score for Hard difficulty: %d attempts!\n", highScores["hard"])
			}
			return
		} else if guessNo < computerNo {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guessNo)
		} else {
			fmt.Printf("Incorrect! The number is less than %d.\n", guessNo)
		}
	}
	fmt.Printf("\nSorry! You've run out of attempts. The correct number was %d.\n", computerNo)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	highScores := map[string]int{
		"easy":   0,
		"medium": 0,
		"hard":   0,
	}

	for {
		RulesOfGame()
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			EasyGame(scanner, highScores)
		case "2":
			MediumGame(scanner, highScores)
		case "3":
			HardGame(scanner, highScores)
		default:
			fmt.Println("Invalid choice. Please select a valid difficulty level (1-3).")
			continue
		}

		fmt.Println("\nCurrent High Scores:")
		fmt.Printf("Easy: %d attempts\n", highScores["easy"])
		fmt.Printf("Medium: %d attempts\n", highScores["medium"])
		fmt.Printf("Hard: %d attempts\n", highScores["hard"])

		if !playAgain() {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}
}
