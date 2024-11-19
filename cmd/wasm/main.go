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

	Arena.DimensionCol = 5
	Arena.DimensionRaw = 5
	js.Global().Call("Arena", Arena.DimensionRaw, Arena.DimensionCol)

	TakedimensionArena()

	return nil
}

func TakedimensionArena() {

	perimeterAre := doc.Call("getElementById", "arena00")
	Arena.Perim.Xleft = int(perimeterAre.Call("getBoundingClientRect").Get("left").Float())
	Arena.Perim.Ytop = int(perimeterAre.Call("getBoundingClientRect").Get("top").Float())

	perimeterAre = doc.Call("getElementById", "arena0"+strconv.Itoa(Arena.DimensionCol-1))
	Arena.Perim.Xright = int(perimeterAre.Call("getBoundingClientRect").Get("right").Float())

	perimeterAre = doc.Call("getElementById", "arena"+strconv.Itoa(Arena.DimensionRaw-1)+strconv.Itoa(Arena.DimensionRaw-1))
	Arena.Perim.Ybottom = int(perimeterAre.Call("getBoundingClientRect").Get("bottom").Float())

}

func CreateTheHero(this js.Value, args []js.Value) interface{} {

	Hero.Position.Xleft = Arena.Perim.Xleft - 40
	Hero.Position.Ytop = Arena.Perim.Ytop

	Hero.Style = Hero.Hero.Get("style")

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.Xleft) + "px"))
	Hero.Style.Call("setProperty", "top", (strconv.Itoa(Hero.Position.Ytop) + "px"))
	js.Global().Call("Hero", Hero.Ctxh)

	return nil
}

func MoveHero(this js.Value, args []js.Value) interface{} {

	if int(Hero.Hero.Call("getBoundingClientRect").Get("right").Float())-20 > Arena.Perim.Xright {

		Hero.Position.Xleft = Arena.Perim.Xleft - 40
		Hero.Position.Ytop += 54
		Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.Xleft) + "px"))
		Hero.Style.Call("setProperty", "top", (strconv.Itoa(Hero.Position.Ytop) + "px"))

		return nil
	}

	Hero.Position.Xleft += 10

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.Xleft) + "px"))

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
