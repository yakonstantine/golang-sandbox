package inf

import "golangbase/tasks/hangman_online/srv/core/ent"

type HangmanRepo interface {
	Add(h *ent.Hangman) (ent.Hangman, error)
	GetByID(id ent.ID) (ent.Hangman, error)
	Update(h *ent.Hangman) error
}
