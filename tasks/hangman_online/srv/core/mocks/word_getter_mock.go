package mocks

import (
	"github.com/stretchr/testify/mock"
	"golangbase/tasks/hangman_online/srv/core/ent"
)

type WordGetterMock struct {
	mock.Mock
}

func (wsm *WordGetterMock) GetWord() (ent.Word, error) {
	args := wsm.Called()
	return args.Get(0).(ent.Word), args.Error(1)
}
