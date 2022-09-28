package guessnumber

import (
	"fmt"
	"math/rand"
	"time"
)

func StartGame(tries int) {
	fmt.Printf("Try guess the number from 1 to 20. You have only %d tries.\n", tries)

	targetNumber := getRandomNumber(1, 20)

	for i := 0; i < tries; i++ {
		if i > 0 {
			fmt.Printf("Try one more time.\n")
		}
		rightAnswer, err := playRound(targetNumber)
		if err != nil {
			fmt.Println(err)
			return
		}
		if rightAnswer {
			break
		}
	}

	fmt.Println("The right answer was", targetNumber)
}

func playRound(targetNumber int) (bool, error) {
	fmt.Print("Enter your guess: ")

	n, err := readFromStdin()
	if err != nil {
		return false, fmt.Errorf("wrong input, the value should be integer, break the game")
	}

	if targetNumber == n {
		fmt.Println("Good job!")
		return true, nil
	}

	var prompt string
	if n > targetNumber {
		prompt = "the target number is less then yours"
	} else {
		prompt = "the target number is bigger then yours"
	}

	fmt.Printf("No, %s. ", prompt)

	return false, nil
}

func getRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func readFromStdin() (n int, err error) {
	_, err = fmt.Scanf("%d\n", &n)
	return
}
