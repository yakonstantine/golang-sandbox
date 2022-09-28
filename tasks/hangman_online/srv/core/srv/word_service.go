package srv

import (
	"context"
	"errors"
	"golangbase/tasks/hangman_online/srv/core/ent"
	"golangbase/tasks/hangman_online/srv/core/inf"
	"time"
)

type WordService struct {
	cancel     func()
	wordBuff   chan ent.Word
	wordGetter inf.WordGetter
}

func NewWordService(parentCtx context.Context, wordGetter inf.WordGetter, buffSize int) *WordService {
	if buffSize <= 0 {
		buffSize = 5
	}
	ctx, cancel := context.WithCancel(parentCtx)
	ws := WordService{
		cancel:     cancel,
		wordBuff:   make(chan ent.Word, buffSize),
		wordGetter: wordGetter,
	}
	go ws.getWords(ctx)
	return &ws
}

func (ws *WordService) GetWord() (ent.Word, error) {
	w, ok := <-ws.wordBuff
	if !ok {
		return "", errors.New("can't return any word")
	}
	return w, nil
}

func (ws *WordService) Cancel() {
	ws.cancel()
}

func (ws *WordService) getWords(ctx context.Context) {
	defer close(ws.wordBuff)
	w, err := ws.wordGetter.GetWord()
	if err != nil {
		return
	}
	for {
		select {
		case <-ctx.Done():
			return
		case ws.wordBuff <- w:
			w, err = ws.wordGetter.GetWord()
			if err != nil {
				return
			}
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
