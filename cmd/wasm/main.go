package main

import (
	"fmt"
	"syscall/js"
)

func CreateTheArena(this js.Value, args []js.Value) interface{} {
	js.Global().Call("Arena", 5, 5)

	return 5
}

func main() {

	fmt.Println("hello")

	js.Global().Set("CreateTheArena", js.FuncOf(CreateTheArena))

	<-make(chan struct{})
}
