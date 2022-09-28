package hangman

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var gameDictionary = map[int]string{
	0: "foo",
	1: "world",
	2: "global",
	3: "door",
}

func StartGame() {
	n := getRandomNumber(len(gameDictionary))
	targetWord := gameDictionary[n]
	answer := createAnswerDraft(len(targetWord))
	misses := make([]string, 0)

	for i := 8; i > 0; {
		printlnAnswer(answer, misses)
		s, err := readFromStdin()
		if err != nil {
			i--
			fmt.Printf("Input error: '%s', '%d' tries left\n", err.Error(), i)
			continue
		}

		if !showLetters(targetWord, s, answer) {
			misses = append(misses, s)
			i--
			fmt.Printf("You missed, '%d' tries left\n", i)
			continue
		}

		if !contains(answer, "_") {
			fmt.Printf("You won! The word is '%s'", targetWord)
			break
		}
	}
}

func showLetters(targetWord, subS string, answer []string) bool {
	res := false
	i := 0
	for i < len(targetWord) {
		idx := strings.Index(targetWord[i:], subS)
		if idx < 0 {
			return res
		}

		i = idx + i
		res = true
		answer[i] = subS
		i++
	}

	return res
}

func getRandomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func createAnswerDraft(wordLen int) []string {
	answer := make([]string, wordLen)
	for i := 0; i < wordLen; i++ {
		answer[i] = "_"
	}

	return answer
}

func printlnAnswer(answer []string, misses []string) {
	fmt.Printf("Word: %s\nMisses: %s\nGuess: ",
		strings.Join(answer, " "), strings.Join(misses, ","))
}

func readFromStdin() (str string, err error) {
	_, err = fmt.Scanf("%s\n", &str)
	if len(str) > 1 || len(str) == 0 {
		return str, errors.New("the answer should cointain only one symbol.")
	}
	return
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
