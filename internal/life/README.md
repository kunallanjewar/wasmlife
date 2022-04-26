# life
--
    import "github.com/kunallanjewar/wasmlife/internal/life"

life manages the state of cells based on B3/S23 rule of Conways Game of Live.

## Usage

```go
var (
	PixelColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
)
```

#### func  ApplyRuleB3S23

```go
func ApplyRuleB3S23(
	life Life,
	current, next *cel.Cell,
	neighbors int) bool
```
ApplyRuleB3S23 applies Conway's Game Of Life rule as following

    1. If a dead cell had exactly 3 alive neighbors, it becomes alive.
    2. If an alive cell had less than 2 or more than 3 alive neighbors, it becomes dead.
    3. If an alive cell has 2 or 3 neighbors, its remains alive.

#### type Life

```go
type Life map[cel.Cell]bool
```

Life is map of living cells.

#### func  New

```go
func New() *Life
```
NewLife instantiates Life. You much call Seed() to populate it later.

#### func (Life) Draw

```go
func (l Life) Draw(screen *ebiten.Image)
```
Draw is called every tick and draws life's next state on screen.

#### func (Life) Seed

```go
func (l Life) Seed(cells []*cel.Cell)
```
Seed populates Life given a slice of living cells.

#### func (Life) Update

```go
func (l Life) Update() error
```
Update is called every tick and generates life's next state.
