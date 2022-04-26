package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kunallanjewar/wasmlife/internal/controller"
	"github.com/kunallanjewar/wasmlife/internal/game"
	"github.com/kunallanjewar/wasmlife/internal/world"
)

const (
	Title = "Game Of Life - Kunal Lanjewar"
)

func main() {
	// game settings
	ebiten.SetWindowSize(game.Width, game.Height)
	ebiten.SetWindowTitle(Title)

	// init deps
	ct := controller.New()
	w := world.New(ct)

	// start game
	err := ebiten.RunGame(game.New(w, ct))
	if err != nil {
		log.Fatal(err)
	}
}
