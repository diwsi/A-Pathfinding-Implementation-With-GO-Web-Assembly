set GOARCH=wasm
set GOOS=js
go build -o Web/pathfinder.wasm main.go
