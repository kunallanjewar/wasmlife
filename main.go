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
	ebiten.SetMaxTPS(30)

	// init deps
	ct := controller.New()
	w := world.New(ct)

	// start game
	err := ebiten.RunGame(
		&game.Game{
			Controller: ct,
			World:      w,
		})

	if err != nil {
		log.Fatal(err)
	}
}
