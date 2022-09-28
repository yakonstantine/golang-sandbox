package ent

const BlankRune rune = '_'

type Hangman struct {
	ID       ID
	Word     []rune
	State    []rune
	Misses   []rune
	TriesCnt int
	Won      bool
}

func NewHangman(w Word) *Hangman {
	wordRunes := []rune(w)
	wordLen := len(wordRunes)
	currentState := make([]rune, wordLen)
	for i := 0; i < wordLen; i++ {
		currentState[i] = BlankRune
	}
	return &Hangman{
		Word:     wordRunes,
		State:    currentState,
		TriesCnt: 8,
		Won:      false,
	}
}
