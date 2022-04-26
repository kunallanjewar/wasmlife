package life

import (
	cel "github.com/kunallanjewar/wasmlife/internal/cell"
)

// ApplyRuleB3S23 applies Conway's Game Of Life rule as following
//
//	1. If a dead cell had exactly 3 alive neighbors, it becomes alive.
//	2. If an alive cell had less than 2 or more than 3 alive neighbors, it becomes dead.
//	3. If an alive cell has 2 or 3 neighbors, its remains alive.
func ApplyRuleB3S23(
	life Life,
	current, next *cel.Cell,
	neighbors int) bool {

	_, alive := life[*current]

	// 1. BIRTH
	if neighbors == 3 && !alive {
		next.Alive(true)
		return true
	}

	// 2. DEATH
	if (neighbors < 2 || neighbors > 3) && alive {
		next.Alive(false)
		return true
	}

	// 3. SURVIVAL
	return false
}
