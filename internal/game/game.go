// game object is an object that needs to be updated and drawn on every frame.
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

type Game struct {
	World      *world.World
	Controller *controller.Controller
}

func (g *Game) Update() error {
	g.Controller.Update()
	g.World.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Controller.Draw(screen)
	g.World.Draw(screen)

	// print fps on top left corner
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("FPS: %0.2f\n", ebiten.CurrentFPS()),
	)
}

func (g *Game) Layout(outW, outH int) (w, h int) {
	return Width / 2, Height / 2
}
