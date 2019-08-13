# A* Pathfinding Implementation With GO Web Assembly
An experimental WASM project that written in GOlang. WASM application solves path problem by using A* pathfinding algorithm.Matrix dataset created by javascript passes to WASM application. Solution returns to JS code and renders on HTML canvas.

To Run Online: <br>
https://diwsi.github.io/index.html
<br>
<br>
To Run Locale:
<br><b>go run .\web\server.go</b><br>
<i>By default: http://localhost:8080/Web/</i>
<br>
<br>
To Build Changes
<br> <b> set GOARCH=wasm</b>
<br>  <b>set GOOS=js</b>
<br> <b> go build -o Web/pathfinder.wasm main.go</b>

