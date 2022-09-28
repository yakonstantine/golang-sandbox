package res

import "golangbase/tasks/hangman_online/srv/core/ent"

type Hangman struct {
	ID       int
	Word     string
	State    string
	Misses   []string
	TriesCnt int
	Won      bool
}

func NewHangman(h *ent.Hangman) *Hangman {
	return &Hangman{
		ID:       int(h.ID),
		Word:     string(h.Word),
		State:    string(h.State),
		Misses:   convertMisses(h.Misses),
		TriesCnt: h.TriesCnt,
		Won:      h.Won,
	}
}

func convertMisses(runes []rune) []string {
	if runes == nil {
		return nil
	}
	res := make([]string, 0, len(runes))
	for _, r := range runes {
		res = append(res, string(r))
	}
	return res
}
