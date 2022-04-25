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

	welcomeScreenTextFmt = `
-----------------
Enter Live Cells:
- Each cell must be separated by a newline char or space.
- Ex. (X1,Y1) (X2, Y2) ...
- Hit TAB once to begin or
- Hit ESC once to accept default (r-pintomino) pattern.
--------
(100,60)
(101,60)
(100,61)
(99,61)
(100,62)
--------
%s
`
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

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) ||
		repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		in.text += "\n"
	}

	// If the backspace key is pressed, remove one character.
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
	if repeatingKeyPressed(ebiten.KeyEscape) {
		in.enabled = false
		in.text = PatternRpentomino
		return nil
	}

	in.counter++

	return nil
}

func (in *Controller) Draw(screen *ebiten.Image) {
	if !in.enabled {
		return
	}

	// Blink the cursor, last char.
	t := in.text
	if in.counter%60 < 30 {
		t += "_"
	}

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(welcomeScreenTextFmt, t),
	)
}

// Input return game input in Cell format.
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

		x, err := strconv.ParseInt(xy[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.ParseInt(xy[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		cells = append(cells, cel.New(x, y))
	}
	return cells
}