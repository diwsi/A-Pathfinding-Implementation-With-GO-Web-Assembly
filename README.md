# A Pathfinding Implementation With GO Web Assembly
An experimental WASM project that written in GOlang. WASM application solves path problem by using A* pathfinding algorithm.Matrix dataset created by javascript passes to WASM application. Solution returns to JS code and renders on HTML canvas.

To run Sample:
# go run .\web\server.go

To Build Changes
# set GOARCH=wasm
# set GOOS=js
# go build -o Web/pathfinder.wasm main.go

