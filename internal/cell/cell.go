package cell

// A Cell is represented as X and Y coordinates in a 2d plane.
type Cell struct {
	x, y  int64
	alive bool
}

func New(x, y int64) *Cell {
	return &Cell{
		x:     x,
		y:     y,
		alive: true,
	}
}

// Alive returns true if this cell alive.
func (c *Cell) IsAlive() bool {
	return c.alive
}

// Alive sets whether cell is dead or alive.
func (c *Cell) Alive(b bool) {
	c.alive = b
}

// OriginalState state return a cell with original living state.
func (c *Cell) OriginalState() *Cell {
	c.alive = true
	return c
}

func (c *Cell) XY() (x, y int64) {
	return c.x, c.y
}

func (c *Cell) X() int64 {
	return c.x
}

func (c *Cell) Y() int64 {
	return c.y
}
