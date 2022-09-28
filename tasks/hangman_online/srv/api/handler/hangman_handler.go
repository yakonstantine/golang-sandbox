package handler

import (
	"fmt"
	"golangbase/tasks/hangman_online/srv/api/res"
	"golangbase/tasks/hangman_online/srv/core/ent"
	"golangbase/tasks/hangman_online/srv/core/inf"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
)

type HangnamHandler struct {
	hangmanService inf.HangmanService
}

func NewHangmanHandler(hs inf.HangmanService) *HangnamHandler {
	return &HangnamHandler{hangmanService: hs}
}

func (hh *HangnamHandler) GetNewHangman(c *fiber.Ctx) error {
	h, err := hh.hangmanService.CreateGame()
	if err != nil {
		return err
	}

	return hangmanToJson(c, h)
}

func (hh *HangnamHandler) GetHangman(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return badRequest(c, err)
	}

	h, err := hh.hangmanService.GetGame(ent.ID(id))
	if err != nil {
		return badRequest(c, err)
	}

	return hangmanToJson(c, h)
}

func (hh *HangnamHandler) PostTryGuess(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return badRequest(c, err)
	}

	body := c.Body()
	r, _ := utf8.DecodeRune(body)

	h, err := hh.hangmanService.TryGuess(ent.ID(id), r)
	if err != nil {
		return badRequest(c, err)
	}

	return hangmanToJson(c, h)
}

func badRequest(c *fiber.Ctx, err error) error {
	return c.Status(400).SendString(fmt.Sprintf("Bad Request: %s", err))
}

func hangmanToJson(c *fiber.Ctx, h *ent.Hangman) error {
	return c.JSON(res.NewHangman(h))
}
