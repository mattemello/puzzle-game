package main

import (
	"fmt"
	strctVar "github.com/mattemello/puzzle-game/cmd/wasm/structAndVar"
	"strconv"
	"syscall/js"
)

var Arena strctVar.TheArena

var Hero strctVar.TheHero

func CreateTheArena(this js.Value, args []js.Value) interface{} {

	Arena.Dimension.X = 5
	Arena.Dimension.Y = 5
	js.Global().Call("Arena", Arena.Dimension.X, Arena.Dimension.Y)

	return nil
}

func CreateTheHero(this js.Value, args []js.Value) interface{} {

	Hero.Position.X = 0
	Hero.Position.Y = 8

	Hero.Style = Hero.Hero.Get("style")

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.X) + "px"))
	Hero.Style.Call("setProperty", "top", (strconv.Itoa(Hero.Position.Y) + "px"))
	js.Global().Call("Hero", Hero.Ctxh)

	return nil
}

func MoveHero(this js.Value, args []js.Value) interface{} {

	Hero.Position.X += 10

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.X) + "px"))

	return nil
}

func main() {

	fmt.Println("hello")

	/*
	   theHero.hero = document.getElementById("hero");
	   theHero.ctxh = theHero.hero.getContext("2d");
	*/

	doc := js.Global().Get("document")

	Hero.Hero = doc.Call("getElementById", "hero")
	Hero.Ctxh = Hero.Hero.Call("getContext", "2d")

	js.Global().Set("CreateTheArena", js.FuncOf(CreateTheArena))
	js.Global().Set("CreateTheHero", js.FuncOf(CreateTheHero))
	js.Global().Set("MoveHero", js.FuncOf(MoveHero))

	<-make(chan struct{})
}
