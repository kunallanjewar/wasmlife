// world manages the game environment and it's child game objects.
package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	cel "github.com/kunallanjewar/wasmlife/internal/cell"
	"github.com/kunallanjewar/wasmlife/internal/life"
)

type Controller interface {
	Input() ([]*cel.Cell, bool)
}

// Life represents state of cells.
type World struct {
	// alive is a set of living cells.
	contrl Controller
	life   *life.Life
	seeded bool
}

// New creates a new instance of World object.
func New(c Controller) *World {
	w := &World{
		contrl: c,
		life:   life.New(),
		seeded: false,
	}
	return w
}

func (w *World) Update() error {
	if !w.seeded {
		w.seed()
		return nil
	}

	return w.life.Update()
}

func (w *World) Draw(screen *ebiten.Image) {
	if !w.seeded {
		return
	}

	w.life.Draw(screen)
}

func (w *World) seed() {
	if !w.seeded {
		cells, ok := w.contrl.Input()
		if ok {
			w.life.Seed(cells)
			w.seeded = true
		}
	}
}
