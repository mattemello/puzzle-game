package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

type Coordination struct {
	x int
	y int
}

type theArena struct {
	Dimension Coordination
}

type theHero struct {
	Hero     js.Value
	Ctxh     js.Value
	Style    js.Value
	Position Coordination
	//Dimension Coordination
}

var Arena theArena

var Hero theHero

func CreateTheArena(this js.Value, args []js.Value) interface{} {

	Arena.Dimension.x = 5
	Arena.Dimension.y = 5
	js.Global().Call("Arena", Arena.Dimension.x, Arena.Dimension.y)

	return nil
}

func CreateTheHero(this js.Value, args []js.Value) interface{} {

	Hero.Position.x = 0
	Hero.Position.y = 8

	Hero.Style = Hero.Hero.Get("style")

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.x) + "px"))
	Hero.Style.Call("setProperty", "top", (strconv.Itoa(Hero.Position.y) + "px"))
	js.Global().Call("Hero", Hero.Ctxh)

	return nil
}

func MoveHero(this js.Value, args []js.Value) interface{} {

	Hero.Position.x += 10

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.x) + "px"))

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