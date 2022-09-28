package inf

import (
	"golangbase/tasks/hangman_online/srv/core/ent"
)

type WordService interface {
	GetWord() (ent.Word, error)
	Cancel()
}
