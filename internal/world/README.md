# world
--
    import "github.com/kunallanjewar/wasmlife/internal/world"

world manages the game environment and it's child game objects.

## Usage

#### type Controller

```go
type Controller interface {
	Input() ([]*cel.Cell, bool)
}
```

Controller is an interface to get game's input.

#### type World

```go
type World struct {
}
```

Life represents state of cells.

#### func  New

```go
func New(c Controller) *World
```
New creates a new instance of World object.

#### func (*World) Draw

```go
func (w *World) Draw(screen *ebiten.Image)
```
Draw is called every frame.

#### func (*World) Update

```go
func (w *World) Update() error
```
Update is called every tick.
