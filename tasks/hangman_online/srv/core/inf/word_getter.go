package inf

import (
	"golangbase/tasks/hangman_online/srv/core/ent"
)

type WordGetter interface {
	GetWord() (ent.Word, error)
}
