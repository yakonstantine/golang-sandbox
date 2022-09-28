package infra

import (
	"fmt"
	"golangbase/tasks/hangman_online/srv/core/ent"
)

type HangmanRepo struct {
	storage map[ent.ID]*ent.Hangman
	nextID  int
}

func NewHangmanRepo() *HangmanRepo {
	return &HangmanRepo{
		storage: make(map[ent.ID]*ent.Hangman),
		nextID:  1,
	}
}

func (hr *HangmanRepo) Add(h *ent.Hangman) (ent.Hangman, error) {
	newH := *h
	newH.ID = ent.ID(hr.nextID)
	hr.storage[newH.ID] = &newH
	hr.nextID++
	return newH, nil
}

func (hr *HangmanRepo) GetByID(id ent.ID) (ent.Hangman, error) {
	h, ok := hr.storage[id]
	if !ok {
		return ent.Hangman{}, fmt.Errorf("hangman game with ID:'%d' doesn't exist", id)
	}
	return *h, nil
}

func (hr *HangmanRepo) Update(h *ent.Hangman) error {
	hr.storage[h.ID] = h
	return nil
}
