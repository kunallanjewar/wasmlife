# cell
--
    import "github.com/kunallanjewar/wasmlife/internal/cell"


## Usage

#### type Cell

```go
type Cell struct {
}
```

A Cell is represented as X and Y coordinates in a 2d plane.

#### func  New

```go
func New(x, y int64) *Cell
```

#### func (*Cell) Alive

```go
func (c *Cell) Alive(b bool)
```
Alive sets whether cell is dead or alive.

#### func (*Cell) IsAlive

```go
func (c *Cell) IsAlive() bool
```
Alive returns true if this cell alive.

#### func (*Cell) OriginalState

```go
func (c *Cell) OriginalState() *Cell
```
OriginalState state return a cell with original living state.

#### func (*Cell) X

```go
func (c *Cell) X() int64
```

#### func (*Cell) XY

```go
func (c *Cell) XY() (x, y int64)
```

#### func (*Cell) Y

```go
func (c *Cell) Y() int64
```
