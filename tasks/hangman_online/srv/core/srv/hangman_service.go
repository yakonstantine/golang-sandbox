package srv

import (
	"errors"
	"golangbase/tasks/hangman_online/srv/core/ent"
	"golangbase/tasks/hangman_online/srv/core/inf"
)

type HangmanService struct {
	wordsService inf.WordService
	hangmanRepo  inf.HangmanRepo
}

func NewHangmanService(ws inf.WordService, hr inf.HangmanRepo) *HangmanService {
	return &HangmanService{
		wordsService: ws,
		hangmanRepo:  hr,
	}
}

func (hs *HangmanService) CreateGame() (*ent.Hangman, error) {
	word, err := hs.wordsService.GetWord()
	if err != nil {
		return nil, err
	}
	h := ent.NewHangman(word)
	savedH, err := hs.hangmanRepo.Add(h)
	if err != nil {
		return nil, err
	}
	return &savedH, nil
}

func (hs *HangmanService) GetGame(id ent.ID) (*ent.Hangman, error) {
	h, err := hs.hangmanRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (hs *HangmanService) TryGuess(id ent.ID, r rune) (h *ent.Hangman, err error) {
	saveH, err := hs.hangmanRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if saveH.TriesCnt <= 0 {
		return nil, errors.New("all tries have been used")
	}
	if saveH.Won {
		return nil, errors.New("the game has been won")
	}
	defer func() {
		err = hs.hangmanRepo.Update(h)
		if err != nil {
			h = nil
		}
	}()
	h = &saveH
	success := false
	for i, rv := range h.Word {
		if rv != r {
			continue
		}
		h.State[i] = r
		success = true
	}
	if !success {
		h.TriesCnt--
		h.Misses = append(h.Misses, r)
	}
	if success && !contains(h.State, ent.BlankRune) {
		h.Won = true
	}
	return
}

func contains(runes []rune, r rune) bool {
	for _, rv := range runes {
		if rv == r {
			return true
		}
	}
	return false
}
