# controller
--
    import "github.com/kunallanjewar/wasmlife/internal/controller"


## Usage

```go
const (
	/*
	   -------------
	   ------X------
	   ------X------
	   ------X------
	   -------------
	*/
	PatternOscillator = `
(125,125)
(125,124)
(125,123)
`
	/*
	   -------------
	   ------X-X----
	   ----X-X------
	   ------X------
	   -------------
	*/
	PatternRpentomino = `
(150,150)
(151,152)
(150,151)
(149,151)
(150,152)
`
	/*
	   --------------
	   ---------X----
	   -----X---X----
	   -------X-X----
	   --------------
	*/
	PatternGlider = `
(150,150)
(151,149)
(152,149)
(152,150)
(152,151)
`
)
```

#### type Controller

```go
type Controller struct {
}
```

Controller provides game's input.

#### func  New

```go
func New() *Controller
```
New creates an instrance of a Controller.

#### func (*Controller) Draw

```go
func (in *Controller) Draw(screen *ebiten.Image)
```
Update run every frame and draws on screen.

#### func (*Controller) Input

```go
func (in *Controller) Input() ([]*cel.Cell, bool)
```
Input returns game input in Cell format.

#### func (*Controller) Update

```go
func (in *Controller) Update() error
```
Update run every tick and check for user input.
