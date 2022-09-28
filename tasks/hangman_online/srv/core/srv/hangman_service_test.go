package srv

import (
	"errors"
	"golangbase/tasks/hangman_online/srv/core/ent"
	"golangbase/tasks/hangman_online/srv/core/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHangmanService_CreateGame(t *testing.T) {
	// arrange
	word := ent.Word("foo")
	savedHangman := ent.NewHangman(word)

	wsm := new(mocks.WordServiceMock)
	wsm.On("GetWord").Return(word, nil)
	hrm := new(mocks.HangmanRepoMock)
	hrm.On("Add", mock.Anything).Return(*savedHangman, nil)

	sut := NewHangmanService(wsm, hrm)
	// act
	res, err := sut.CreateGame()
	// assert
	if assert.Nil(t, err) {
		assert.Equal(t, res, savedHangman)
	}
}

func TestHangmanService_CreateGame_ReturnsError_WordsServiceError(t *testing.T) {
	// arrange
	expErr := errors.New("error")

	wsm := new(mocks.WordServiceMock)
	wsm.On("GetWord").Return(ent.Word(""), expErr)
	hrm := new(mocks.HangmanRepoMock)

	sut := NewHangmanService(wsm, hrm)
	// act
	_, err := sut.CreateGame()
	// assert
	if assert.Error(t, err) {
		assert.Equal(t, expErr, err)
	}
}

func TestHangmanService_CreateGame_ReturnsError_HangmanRepoError(t *testing.T) {
	// arrange
	word := ent.Word("foo")
	expErr := errors.New("error")

	wsm := new(mocks.WordServiceMock)
	wsm.On("GetWord").Return(word, nil)
	hrm := new(mocks.HangmanRepoMock)
	hrm.On("Add", mock.Anything).Return(ent.Hangman{}, expErr)

	sut := NewHangmanService(wsm, hrm)
	// act
	_, err := sut.CreateGame()
	// assert
	if assert.Error(t, err) {
		assert.Equal(t, expErr, err)
	}
}

func TestHangmanService_GetGame(t *testing.T) {
	// arrange
	word := ent.Word("foo")
	id := ent.ID(1)
	savedHangman := ent.NewHangman(word)
	savedHangman.ID = id

	wsm := new(mocks.WordServiceMock)
	hrm := new(mocks.HangmanRepoMock)
	hrm.On("GetByID", id).Return(*savedHangman, nil)

	sut := NewHangmanService(wsm, hrm)
	// act
	res, err := sut.GetGame(id)
	// assert
	if assert.Nil(t, err) {
		assert.Equal(t, res, savedHangman)
	}
}

func TestHangmanService_GetGame_ReturnsError_HangmanRepoError(t *testing.T) {
	// arrange
	id := ent.ID(1)
	expErr := errors.New("error")

	wsm := new(mocks.WordServiceMock)
	hrm := new(mocks.HangmanRepoMock)
	hrm.On("GetByID", id).Return(ent.Hangman{}, expErr)

	sut := NewHangmanService(wsm, hrm)
	// act
	_, err := sut.GetGame(id)
	// assert
	if assert.Error(t, err) {
		assert.Equal(t, expErr, err)
	}
}

func TestHangmanService_TryGuess_OneRune(t *testing.T) {
	// arrange
	hangman := ent.NewHangman("foo")
	hangman.ID = ent.ID(1)
	expState := ent.Word("f__")

	sut := setUpSUT(hangman)
	// act
	res, err := sut.TryGuess(hangman.ID, 'f')
	// assert
	if !assert.Nil(t, err) {
		return
	}
	actState := ent.Word(res.State)
	assert.Equal(t, expState, actState)
	assert.Equal(t, 0, len(res.Misses))
}

func TestHangmanService_TryGuess_TwoRunes(t *testing.T) {
	// arrange
	hangman := ent.NewHangman("foo")
	hangman.ID = ent.ID(1)
	expState := ent.Word("_oo")

	sut := setUpSUT(hangman)
	// act
	res, err := sut.TryGuess(hangman.ID, 'o')
	// assert
	if !assert.Nil(t, err) {
		return
	}
	actState := ent.Word(res.State)
	assert.Equal(t, expState, actState)
	assert.Equal(t, 0, len(res.Misses))
}

func TestHangmanService_TryGuess_WrongRune(t *testing.T) {
	// arrange
	hangman := ent.NewHangman("foo")
	hangman.ID = ent.ID(1)
	expState := ent.Word("___")
	tryRune := 'g'

	sut := setUpSUT(hangman)
	// act
	res, err := sut.TryGuess(hangman.ID, tryRune)
	// assert
	if !assert.Nil(t, err) {
		return
	}
	actState := ent.Word(res.State)
	assert.Equal(t, expState, actState)
	assert.Equal(t, 1, len(res.Misses))
	if !contains(res.Misses, tryRune) {
		assert.Failf(t, "Fail", "Expected '%s' but returned '%v'", string(tryRune), string(res.Misses))
	}
}

func TestHangmanService_TryGuess_WonTheGame(t *testing.T) {
	// arrange
	expMisses := []rune{'a', 'b', 'c'}
	hangman := ent.NewHangman("foo")
	hangman.ID = ent.ID(1)
	hangman.State = []rune("_oo")
	hangman.Misses = append(hangman.Misses, expMisses...)
	hangman.TriesCnt -= len(expMisses)
	tryRune := 'f'

	sut := setUpSUT(hangman)
	// act
	res, err := sut.TryGuess(hangman.ID, tryRune)
	// assert
	if !assert.Nil(t, err) {
		return
	}

	assert.True(t, res.Won)
	expState := ent.Word("foo")
	actState := ent.Word(res.State)
	assert.Equal(t, expState, actState)
	assert.Equal(t, len(expMisses), len(res.Misses))
	expMissesStr := string(expMisses)
	actMissesStr := string(res.Misses)
	assert.Equal(t, expMissesStr, actMissesStr)
}

func TestHangmanService_TryGuess_LooseTheGame(t *testing.T) {
	// arrange
	expMisses := []rune{'a', 'b', 'c', 'd', 'e'}
	hangman := ent.NewHangman("foo")
	hangman.ID = ent.ID(1)
	hangman.State = []rune("_oo")
	hangman.Misses = append(hangman.Misses, expMisses...)
	hangman.TriesCnt -= len(expMisses)
	tryRune := 'g'

	sut := setUpSUT(hangman)
	// act
	res, err := sut.TryGuess(hangman.ID, tryRune)
	// assert
	if !assert.Nil(t, err) {
		return
	}

	assert.False(t, res.Won)
	expState := ent.Word("_oo")
	actState := ent.Word(res.State)
	assert.Equal(t, expState, actState)
	assert.Equal(t, len(expMisses)+1, len(res.Misses))
	expMissesStr := string(append(expMisses, tryRune))
	actMissesStr := string(res.Misses)
	assert.Equal(t, expMissesStr, actMissesStr)
}

func TestHangmanService_TryGuess_NoMoreTries_Error(t *testing.T) {
	// arrange
	hangman := ent.NewHangman("foo")
	hangman.ID = ent.ID(1)
	hangman.TriesCnt = 0

	sut := setUpSUT(hangman)
	// act
	_, err := sut.TryGuess(hangman.ID, 'f')
	// assert
	assert.Error(t, err)
}

func TestHangmanService_TryGuess_AlreadyWon_Error(t *testing.T) {
	// arrange
	hangman := ent.NewHangman("foo")
	hangman.ID = ent.ID(1)
	hangman.Won = true

	sut := setUpSUT(hangman)
	// act
	_, err := sut.TryGuess(hangman.ID, 'f')
	// assert
	assert.Error(t, err)
}

func TestHangmanService_TryGuess_NotFoundInRepo_Error(t *testing.T) {
	// arrange
	hangman := ent.NewHangman("foo")
	hangman.ID = ent.ID(1)
	expErr := errors.New("error")

	wsm := new(mocks.WordServiceMock)
	hrm := new(mocks.HangmanRepoMock)
	hrm.On("GetByID", mock.Anything).Return(ent.Hangman{}, expErr)

	sut := NewHangmanService(wsm, hrm)
	// act
	_, err := sut.TryGuess(hangman.ID, 'f')
	// assert
	if assert.Error(t, err) {
		assert.Equal(t, expErr, err)
	}
}

func TestHangmanService_TryGuess_FailDuringUpdate_Error(t *testing.T) {
	// arrange
	hangman := ent.NewHangman("foo")
	hangman.ID = ent.ID(1)
	expErr := errors.New("error")

	wsm := new(mocks.WordServiceMock)
	hrm := new(mocks.HangmanRepoMock)
	hrm.On("GetByID", hangman.ID).Return(*hangman, nil)
	hrm.On("Update", mock.Anything).Return(expErr)

	sut := NewHangmanService(wsm, hrm)
	// act
	_, err := sut.TryGuess(hangman.ID, 'f')
	// assert
	if assert.Error(t, err) {
		assert.Equal(t, expErr, err)
	}
}

func setUpSUT(h *ent.Hangman) *HangmanService {
	wsm := new(mocks.WordServiceMock)
	hrm := new(mocks.HangmanRepoMock)
	hrm.On("GetByID", h.ID).Return(*h, nil)
	hrm.On("Update", mock.Anything).Return(nil)

	return NewHangmanService(wsm, hrm)
}
