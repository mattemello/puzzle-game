package main

import (
	"fmt"
	"syscall/js"
)

func CreateTheArena(this js.Value, args []js.Value) interface{} {
	js.Global().Call("Arena", 5, 5)

	return nil
}

func CreateTheHero(this js.Value, args []js.Value) interface{} {

	Hero := js.Global().Get("theHero")
	Hero.Set("position.x", "0")
	Hero.Set("position.y", "8")

	js.Global().Call("Hero", 5, 5)

	return nil
}

func main() {

	fmt.Println("hello")

	js.Global().Set("CreateTheArena", js.FuncOf(CreateTheArena))
	js.Global().Set("CreateTheHero", js.FuncOf(CreateTheHero))

	<-make(chan struct{})
}
