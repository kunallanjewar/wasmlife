// life manages the state of cells based on B3/S23 rule of Conways Game of Live.
package life

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	cel "github.com/kunallanjewar/wasmlife/internal/cell"
)

var (
	PixelColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
)

// Life is map of living cells.
type Life map[cel.Cell]bool

// NewLife instantiates Life.
// You much call Seed() to populate it later.
func New() *Life {
	m := make(Life)
	return &m
}

// Seed populates Life given a slice of living cells.
func (l Life) Seed(cells []*cel.Cell) {
	for _, cell := range cells {
		l[*cell] = true
	}
}

func (l Life) Update() error {
	l.update()
	return nil
}

func (l Life) Draw(screen *ebiten.Image) {
	for k := range l {
		screen.Set(int(k.X()), int(k.Y()), PixelColor)
	}
}

func (l Life) update() {
	var newlife []*cel.Cell

	for cell := range l.queue() {
		if c := l.nextState(cell); c != nil {
			newlife = append(newlife, c)
		}
	}

	l.rePopulate(newlife)
}

func (l Life) queue() <-chan *cel.Cell {
	ch := make(chan *cel.Cell, 1)

	go func() {
		defer close(ch)

		for k := range l {
			visitNeighbors(
				k.X(), k.Y(),
				func(dx, dy int64) {
					ch <- cel.New(dx, dy)
				},
			)
		}
	}()

	return ch
}

// nextState computes the next state of the current living cells
// based on game rules.
func (l Life) nextState(cc *cel.Cell) *cel.Cell {
	neighbor := 0
	visitNeighbors(
		cc.X(), cc.Y(),
		// if neighbor is alive, then incr total count
		func(dx, dy int64) {
			// skip current cell in total count
			if cc.X() == dx && cc.Y() == dy {
				return
			}

			if ok := l[*cel.New(dx, dy)]; ok {
				neighbor += 1
			}
		},
	)

	// this shallow copy is not necessary but good for redability.
	next := cc

	if ok := ApplyRuleB3S23(l, cc, next, neighbor); ok {
		return next
	}

	return nil
}

func (l Life) rePopulate(cells []*cel.Cell) {
	for _, cell := range cells {
		if cell.IsAlive() {
			l[*cell] = true
			continue
		}

		if _, ok := l[*cell.OriginalState()]; ok {
			delete(l, *cell)
		}
	}
}

// do func executes given a x, y coord.
type do func(x, y int64)

// visitNeighbors visits all 8 neighbors given x, y coordinates of
// current cell and executes some logic wrapped in func Do().
func visitNeighbors(x, y int64, fn do) {
	dir_x := []int64{x - 1, x, x + 1}
	dir_y := []int64{y - 1, y, y + 1}

	for _, dx := range dir_x {
		for _, dy := range dir_y {
			fn(dx, dy)
		}
	}
}
