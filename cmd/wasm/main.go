package main

import (
	"fmt"
	"strconv"
	"syscall/js"

	strctVar "github.com/mattemello/puzzle-game/cmd/wasm/structAndVar"
)

/* NOTE: the approssimative position of the circle hero:
 * right: ~(-40)
 * left: ~(-30)
 * top: ~(-8)
 * bottom: ~(-8)
 */

var Arena strctVar.TheArena

var Hero strctVar.TheHero

var doc = js.Global().Get("document")

func CreateTheArena(this js.Value, args []js.Value) interface{} {

	Arena.Dimension.X = 5
	Arena.Dimension.Y = 5
	js.Global().Call("Arena", Arena.Dimension.X, Arena.Dimension.Y)

	return nil
}

func CreateTheHero(this js.Value, args []js.Value) interface{} {

	startArena := doc.Call("getElementById", "arena00")

	x := startArena.Call("getBoundingClientRect").Get("left").Float()
	y := startArena.Call("getBoundingClientRect").Get("top").Float()

	fmt.Println(startArena)

	fmt.Println(startArena.Call("getBoundingClientRect").Get("left").Float())
	fmt.Println(startArena.Call("getBoundingClientRect").Get("top").Float())
	fmt.Println(startArena.Call("getBoundingClientRect").Get("right").Float())
	fmt.Println(startArena.Call("getBoundingClientRect").Get("bottom").Float())

	Hero.Position.X = int(x) - 40
	Hero.Position.Y = int(y)

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

	Hero.Hero = doc.Call("getElementById", "hero")
	Hero.Ctxh = Hero.Hero.Call("getContext", "2d")

	js.Global().Set("CreateTheArena", js.FuncOf(CreateTheArena))
	js.Global().Set("CreateTheHero", js.FuncOf(CreateTheHero))
	js.Global().Set("MoveHero", js.FuncOf(MoveHero))

	<-make(chan struct{})
}
