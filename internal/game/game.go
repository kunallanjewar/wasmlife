// game is an object that needs to be updated and drawn on every frame.
package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kunallanjewar/wasmlife/internal/controller"
	"github.com/kunallanjewar/wasmlife/internal/world"
)

const (
	Width  = 640
	Height = 480
)

// Game manages the world object and it's controls.
type Game struct {
	world      *world.World
	controller *controller.Controller
}

// New creates an instance of Game object.
func New(w *world.World, c *controller.Controller) *Game {
	return &Game{w, c}
}

// Update is called every tick.
func (g *Game) Update() error {
	g.controller.Update()
	g.world.Update()

	return nil
}

// Draw is called every frame.
func (g *Game) Draw(screen *ebiten.Image) {
	g.controller.Draw(screen)
	g.world.Draw(screen)

	// print fps on top left corner
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("FPS: %0.2f\n", ebiten.CurrentFPS()),
	)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
//
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
// Ref: Ebiten [docs](https://ebiten.org/documents/cheatsheet.html).
func (g *Game) Layout(outW, outH int) (w, h int) {
	return Width / 2, Height / 2
}
