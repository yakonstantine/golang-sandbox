package infra

import (
	"golangbase/tasks/hangman_online/srv/core/ent"
	"io"
	"net/http"
)

type WordsProxy struct {
	url string
}

func NewWordsProxy(url string) *WordsProxy {
	return &WordsProxy{url: url}
}

func (wp *WordsProxy) GetWord() (ent.Word, error) {
	resp, err := http.Post(wp.url, "text/html; charset=utf-8", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return ent.Word(body), nil
}
