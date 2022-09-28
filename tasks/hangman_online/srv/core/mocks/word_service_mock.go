package mocks

import (
	"github.com/stretchr/testify/mock"
	"golangbase/tasks/hangman_online/srv/core/ent"
)

type WordServiceMock struct {
	mock.Mock
}

func (wsm *WordServiceMock) GetWord() (ent.Word, error) {
	args := wsm.Called()
	return args.Get(0).(ent.Word), args.Error(1)
}

func (wsm *WordServiceMock) Cancel() {
	wsm.Called()
}
