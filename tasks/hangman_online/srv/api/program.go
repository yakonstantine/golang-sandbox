package api

import (
	"context"
	"golangbase/tasks/hangman_online/srv/api/handler"
	"golangbase/tasks/hangman_online/srv/core/inf"
	"golangbase/tasks/hangman_online/srv/core/srv"
	"golangbase/tasks/hangman_online/srv/infra"
	"log"

	_ "golangbase/cmd/hangman_online/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()
	api := app.Group("/api")
	api.Get("/swagger/*", swagger.Handler)

	ctx := context.Background()
	wg := infra.NewWordsProxy("http://watchout4snakes.com/Random/RandomWord")
	ws := srv.NewWordService(ctx, wg, 5)

	setupWordHandler(api, ws)
	setupHangmanHandler(api, ws)

	log.Fatal(app.Listen(":44100"))
}

func setupWordHandler(r fiber.Router, ws inf.WordService) {
	wh := handler.NewWordHandler(ws)

	r.Get("/words", wh.GetWord)
}

func setupHangmanHandler(r fiber.Router, ws inf.WordService) {
	hr := infra.NewHangmanRepo()
	hs := srv.NewHangmanService(ws, hr)
	hh := handler.NewHangmanHandler(hs)

	r.Get("/hangman", hh.GetNewHangman)
	r.Get("/hangman/:id", hh.GetHangman)
	r.Post("/hangman/:id/try-guess", hh.PostTryGuess)
}
