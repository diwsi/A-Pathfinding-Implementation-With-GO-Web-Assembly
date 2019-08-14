package main

import (
	"AStar"
	"syscall/js"
)

func Resolve(this js.Value, args []js.Value) interface{} {
	//Dimentions of matrix
	var w = args[0].Int()
	var h = args[1].Int()

	problem := make([][]uint8, h)
	for i := range problem {
		problem[i] = make([]uint8, w)
	}

	//Must be faster way to build matrix. Maybe with CopyBytesToGo in go v1.13
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			problem[y][x] = uint8(args[y*w+x+2].Int())
		}
	}

	solution := AStar.Resolve(problem)
	//Convert to js array
	return js.TypedArrayOf(solution)
}

func main() {
	//Channel to keep app alive
	c := make(chan bool)

	//Call back function for js
	js.Global().Set("Resolve", js.FuncOf(Resolve))

	c <- true
}
