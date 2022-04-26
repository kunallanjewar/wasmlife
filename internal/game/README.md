# game
--
    import "github.com/kunallanjewar/wasmlife/internal/game"

game is an object that needs to be updated and drawn on every frame.

## Usage

```go
const (
	Width  = 640
	Height = 480
)
```

#### type Game

```go
type Game struct {
}
```

Game manages the world object and it's controls.

#### func  New

```go
func New(w *world.World, c *controller.Controller) *Game
```
New creates an instance of Game object.

#### func (*Game) Draw

```go
func (g *Game) Draw(screen *ebiten.Image)
```
Draw is called every frame.

#### func (*Game) Layout

```go
func (g *Game) Layout(outW, outH int) (w, h int)
```
Layout takes the outside size (e.g., the window size) and returns the (logical)
screen size.

If you don't have to adjust the screen size with the outside size, just return a
fixed size. Ref: Ebiten [docs](https://ebiten.org/documents/cheatsheet.html).

#### func (*Game) Update

```go
func (g *Game) Update() error
```
Update is called every tick.
