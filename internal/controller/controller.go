// controller controlls game's input.
package controller

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	cel "github.com/kunallanjewar/wasmlife/internal/cell"
)

const (
	delay    = 30
	interval = 3
)

type Controller struct {
	runes   []rune
	text    string
	counter int
	enabled bool
}

func New() *Controller {
	return &Controller{
		enabled: true,
	}
}

func (in *Controller) Update() error {
	if !in.enabled {
		return nil
	}

	r := ebiten.AppendInputChars(in.runes[:0])
	in.text += string(r)

	// if the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) ||
		repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		in.text += "\n"
	}

	// if the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(in.text) >= 1 {
			in.text = in.text[:len(in.text)-1]
		}
	}

	// close input loop
	if repeatingKeyPressed(ebiten.KeyTab) {
		in.enabled = false
		return nil
	}

	// default input
	if in.handleDefaultInput() {
		in.enabled = false
		return nil
	}

	in.counter++

	return nil
}

func (in *Controller) Draw(screen *ebiten.Image) {
	if !in.enabled {
		return
	}

	// blink the cursor, last char.
	t := in.text
	if in.counter%60 < 30 {
		t += "_"
	}

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(welcomeScreenTextFmt, t),
	)
}

// Input returns game input in Cell format.
func (in *Controller) Input() ([]*cel.Cell, bool) {
	if in.enabled {
		return nil, false
	}

	return in.parseText(), true
}

func (in *Controller) parseText() []*cel.Cell {
	cells := make([]*cel.Cell, 0)

	ss := strings.Fields(in.text)

	for _, s := range ss {
		sp := s[1 : len(s)-1] // strip suffix and prefix
		xy := strings.Split(sp, ",")

		f := func(s string) int64 {
			n, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			return n
		}

		x := f(xy[0])
		y := f(xy[1])

		cells = append(cells, cel.New(x, y))
	}

	return cells
}

func (in *Controller) handleDefaultInput() bool {

	if repeatingKeyPressed(ebiten.KeyR) {
		in.text = PatternRpentomino
		return true
	}

	if repeatingKeyPressed(ebiten.KeyO) {
		in.text = PatternOscillator
		return true
	}

	if repeatingKeyPressed(ebiten.KeyG) {
		in.text = PatternGlider
		return true
	}

	return false
}
