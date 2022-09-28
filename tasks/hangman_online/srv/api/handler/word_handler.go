package handler

import (
	"golangbase/tasks/hangman_online/srv/core/inf"

	"github.com/gofiber/fiber/v2"
)

type WordHandler struct {
	wordService inf.WordService
}

func NewWordHandler(ws inf.WordService) *WordHandler {
	return &WordHandler{wordService: ws}
}

func (wh *WordHandler) GetWord(c *fiber.Ctx) error {
	word, err := wh.wordService.GetWord()
	if err != nil {
		return err
	}

	return c.JSON(word)
}
