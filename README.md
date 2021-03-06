# wasmlife

`wasmlife` is an awesome _Life_ (Conway's Game Of Life) implementation in Go ❤️ WebAssembly.

## Implementation
_Life_ board is represented in a 64-bit 2D space. 

Since living cells are sparse we used hash map representing a sparse matrix to only track living cells. 

On each tick we visit all the neighbors of each cell and compute the next generation.

For visualization, we used [Ebiten](https://ebiten.org/), a simple 2d game library for Go, which also allows us to compile to WebAssembly. 

## Build
[Go](https://go.dev/doc/install) must be installed on your system.

Standard
```
make build
```

WebAssembly
```
make wasm
```

Run without building
```
make run
```

## CICD
_GitHub-Actions_ builds and auto-commits built WebAssembly binary to `/docs` directory.

_GitHub-Pages_ deploys latest binary automatically.

## Demo
Latest WebAssembly build is visible on [GitHub-Pages](https://kunallanjewar.github.io/wasmlife/)

## Improvements
Many improvements can be made to this project for better UI/UX.

Such as:
- Be able to pan Viewport and/or pan to a specific location.
- Better UI design for taking input.
- Be able to reset the game at runtime.

## License
MIT Licensed