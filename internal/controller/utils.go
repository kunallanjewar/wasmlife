package controller

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func repeatingKeyPressed(key ebiten.Key) bool {

	dt := inpututil.KeyPressDuration(key)
	if dt == 1 {
		return true
	}

	if dt >= delay && (dt-delay)%interval == 0 {
		return true
	}

	return false
}
