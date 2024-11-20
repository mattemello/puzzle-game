package main

import (
	"fmt"
	"math/rand"
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

var Screen strctVar.Screen
var Arena strctVar.TheArena
var Hero strctVar.TheHero
var Path []strctVar.Path

var doc = js.Global().Get("document")

func CreateThePath() {
	// NOTE: the dimension of the path

	var dimensionPath int

	for ok := true; ok; ok = (dimensionPath < 1) {
		dimensionPath = rand.Intn(Arena.DimensionRaw * Arena.DimensionCol)
	}

	Path = make([]strctVar.Path, dimensionPath)

	// NOTE: the start of the path
	Path[0].Numer1 = rand.Intn(Arena.DimensionCol)
	Path[0].Numer2 = rand.Intn(Arena.DimensionRaw)

	Path[0].Arena = doc.Call("getElementById", "arena"+strconv.Itoa(Path[0].Numer1)+strconv.Itoa(Path[0].Numer2))

	Path[0].Coordination.Xleft = int(Path[0].Arena.Call("getBoundingClientRect").Get("left").Float())
	Path[0].Coordination.Ytop = int(Path[0].Arena.Call("getBoundingClientRect").Get("top").Float())
}

func CreateTheArena(this js.Value, args []js.Value) interface{} {

	Arena.DimensionCol = 10
	Arena.DimensionRaw = 10
	js.Global().Call("Arena", Arena.DimensionRaw, Arena.DimensionCol)

	TakedimensionArena()
	CreateThePath()

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

	Hero.Position.Xleft = Path[0].Coordination.Xleft - 40
	Hero.Position.Ytop = Path[0].Coordination.Ytop

	Hero.Style = Hero.Hero.Get("style")

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.Xleft) + "px"))
	Hero.Style.Call("setProperty", "top", (strconv.Itoa(Hero.Position.Ytop) + "px"))
	js.Global().Call("Hero", Hero.Ctxh)

	return nil
}

func MoveHeroX(this js.Value, args []js.Value) interface{} {

	if int(Hero.Hero.Call("getBoundingClientRect").Get("right").Float())-20 > Arena.Perim.Xright && int(args[0].Float()) == 1 {
		return nil
	}
	if int(Hero.Hero.Call("getBoundingClientRect").Get("left").Float())+30 < Arena.Perim.Xleft && int(args[0].Float()) == -1 {
		return nil
	}

	Hero.Position.Xleft += 10 * int(args[0].Float())

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.Xleft) + "px"))

	return nil
}

func MoveHeroY(this js.Value, args []js.Value) interface{} {

	if int(Hero.Hero.Call("getBoundingClientRect").Get("bottom").Float()) > Arena.Perim.Ybottom && int(args[0].Float()) == 1 {
		return nil
	}
	if int(Hero.Hero.Call("getBoundingClientRect").Get("top").Float())-1 < Arena.Perim.Ytop && int(args[0].Float()) == -1 {
		return nil
	}

	Hero.Position.Ytop += 10 * int(args[0].Float())

	Hero.Style.Call("setProperty", "top", (strconv.Itoa(Hero.Position.Ytop) + "px"))

	return nil
}

func TakeDimensionScreen() {
	wind := js.Global().Get("window")

	Screen.Height = int(wind.Get("innerHeight").Float())
	Screen.Width = int(wind.Get("innerWidth").Float())

	fmt.Println(Screen.Width, "-", Screen.Height)
}

func main() {

	TakeDimensionScreen()

	Hero.Hero = doc.Call("getElementById", "hero")
	Hero.Ctxh = Hero.Hero.Call("getContext", "2d")

	js.Global().Set("CreateTheArena", js.FuncOf(CreateTheArena))
	js.Global().Set("CreateTheHero", js.FuncOf(CreateTheHero))
	js.Global().Set("MoveHeroX", js.FuncOf(MoveHeroX))
	js.Global().Set("MoveHeroY", js.FuncOf(MoveHeroY))

	<-make(chan struct{})
}
