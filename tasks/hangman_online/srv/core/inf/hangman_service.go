package inf

import (
	"golangbase/tasks/hangman_online/srv/core/ent"
)

type HangmanService interface {
	CreateGame() (*ent.Hangman, error)
	GetGame(id ent.ID) (*ent.Hangman, error)
	TryGuess(id ent.ID, r rune) (*ent.Hangman, error)
}
