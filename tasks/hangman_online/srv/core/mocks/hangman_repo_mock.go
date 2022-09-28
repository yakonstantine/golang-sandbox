package mocks

import (
	"github.com/stretchr/testify/mock"
	"golangbase/tasks/hangman_online/srv/core/ent"
)

type HangmanRepoMock struct {
	mock.Mock
}

func (hrm *HangmanRepoMock) Add(h *ent.Hangman) (ent.Hangman, error) {
	args := hrm.Called(h)
	return args.Get(0).(ent.Hangman), args.Error(1)
}

func (hrm *HangmanRepoMock) GetByID(id ent.ID) (ent.Hangman, error) {
	args := hrm.Called(id)
	return args.Get(0).(ent.Hangman), args.Error(1)
}

func (hrm *HangmanRepoMock) Update(h *ent.Hangman) error {
	args := hrm.Called(h)
	return args.Error(0)
}
