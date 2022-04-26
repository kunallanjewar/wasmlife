# wasmlife

`wasmlife` is an awesome _Life_ (Conway's Game Of Life) implementation in Go ❤️ WebAssembly.

## Implementation
_Life_ board is represented in a 64-bit 2D space. 

Since living cells are sparse we used hash map representing a sparse matrix to only keep track to living cells. 

On each tick we visit all the neighbors of each cell and compute the next generation.

For visualization, we used [Ebiten](https://ebiten.org/), a simple 2d game library for Go, which also allows us to compile to WebAssembly. 

## Build
Go must be installed on your system.

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

## Demo
Latest WebAssembly build is on [GitHub-Pages](https://kunallanjewar.github.io/wasmlife/)

## Improvements
TODO

## License
MIT Licensed