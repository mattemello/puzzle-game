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

/**
 * TODO: make the path black, so that it will create when you walk it
 * maybe its possible to take the element near the hero, save them and say if they are walkable or not
 * so i can color the element walkable
 */

var Screen strctVar.Screen
var Arena strctVar.TheArena
var Hero strctVar.TheHero
var Path strctVar.ControllPath

var doc = js.Global().Get("document")

func chooseThePath(block strctVar.Path, dimensioPathNow int) strctVar.Path {

	possibleDirection := make(map[int]string)

	possibleDirection[0] = "up"
	possibleDirection[1] = "down"
	possibleDirection[2] = "left"
	possibleDirection[3] = "right"

	// NOTE: number2 -> left/right number1 -> up/down
	if block.Numer1 == 0 {
		delete(possibleDirection, 0)
	} else if block.Numer1 == Arena.DimensionCol-1 {
		delete(possibleDirection, 1)
	}

	if block.Numer2 == 0 {
		delete(possibleDirection, 2)
	} else if block.Numer2 == Arena.DimensionRaw-1 {
		delete(possibleDirection, 3)
	}

	var whatPos int
	var decision string
	thereis := false

	for !thereis {
		whatPos = rand.Intn(len(possibleDirection))
		decision, thereis = possibleDirection[whatPos]
	}

	var newBlock strctVar.Path

	fmt.Println(possibleDirection)

	switch decision {

	case "up":
		newBlock.Numer1 = block.Numer1 - 1
		newBlock.Numer2 = block.Numer2
		break
	case "down":
		newBlock.Numer1 = block.Numer1 + 1
		newBlock.Numer2 = block.Numer2

		break
	case "left":
		newBlock.Numer2 = block.Numer2 - 1
		newBlock.Numer1 = block.Numer1

		break
	case "right":
		newBlock.Numer2 = block.Numer2 + 1
		newBlock.Numer1 = block.Numer1

		break

	default:
		js.Global().Call("alert", "Error - Error in the creation of the path, value not good")
		js.Global().Get("console").Call("error", "Error - Error in the creation of the path, value not good")
	}

	newBlock.Arena = doc.Call("getElementById", "arena"+strconv.Itoa(newBlock.Numer1)+strconv.Itoa(newBlock.Numer2))

	newBlock.Coordination.Xleft = int(newBlock.Arena.Call("getBoundingClientRect").Get("left").Float())
	newBlock.Coordination.Ytop = int(newBlock.Arena.Call("getBoundingClientRect").Get("top").Float())

	newBlock.Ctx = newBlock.Arena.Call("getContext", "2d")
	js.Global().Call("colorPath", newBlock.Ctx, "#313244")

	return newBlock
}

func CreateThePath() {

	dimensionPath := 0

	for !(dimensionPath > 10 && dimensionPath < (Arena.DimensionCol*Arena.DimensionRaw)/2) {
		dimensionPath = rand.Intn(Arena.DimensionRaw * Arena.DimensionCol)
	}

	Path.ArrayPath = make([]string, dimensionPath)
	Path.Path = make(map[string]strctVar.Path)

	num1 := rand.Intn(Arena.DimensionCol)
	num2 := rand.Intn(Arena.DimensionRaw)

	aren := doc.Call("getElementById", "arena"+strconv.Itoa(num1)+strconv.Itoa(num2))

	var coo strctVar.Coordination
	coo.Xleft = int(aren.Call("getBoundingClientRect").Get("left").Float())
	coo.Ytop = int(aren.Call("getBoundingClientRect").Get("top").Float())

	ctx := aren.Call("getContext", "2d")
	js.Global().Call("colorPath", ctx, "#313244")

	Path.ArrayPath[0] = strconv.Itoa(num1) + strconv.Itoa(num2)
	//Path.Path[Path.ArrayPath[0]] = make(strctVar.Path)
	Path.Path[Path.ArrayPath[0]] = strctVar.Path{Numer1: num1, Numer2: num2, Arena: aren, Ctx: ctx, Coordination: coo}

	for i := 1; i < dimensionPath; i++ {
		block := chooseThePath(Path.Path[Path.ArrayPath[i]], i)
		Path.ArrayPath[i] = strconv.Itoa(block.Numer1) + strconv.Itoa(block.Numer2)
		Path.Path[Path.ArrayPath[i]] = block
	}

	for key, ele := range Path.Path {
		fmt.Println("value key "+key+" - ", ele)
	}

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

	Hero.Position.Xleft = Path.Path[Path.ArrayPath[0]].Coordination.Xleft - 40
	Hero.Position.Ytop = Path.Path[Path.ArrayPath[0]].Coordination.Ytop

	Hero.PathCurrentIn.Num1 = Path.Path[Path.ArrayPath[0]].Numer1
	Hero.PathCurrentIn.Num2 = Path.Path[Path.ArrayPath[0]].Numer2

	Hero.Style = Hero.Hero.Get("style")

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.Xleft) + "px"))
	Hero.Style.Call("setProperty", "top", (strconv.Itoa(Hero.Position.Ytop) + "px"))
	js.Global().Call("Hero", Hero.Ctxh)

	return nil
}

func MoveHeroX(this js.Value, args []js.Value) interface{} {

	if int(args[0].Float()) == 1 {
		if int(Hero.Hero.Call("getBoundingClientRect").Get("right").Float()) > Arena.Perim.Ybottom {
			return nil
		}

		if _, ok := Path.Path[strconv.Itoa(Hero.PathCurrentIn.Num1)+strconv.Itoa(Hero.PathCurrentIn.Num2+1)]; !ok {
			return nil
		}
	}
	if int(args[0].Float()) == -1 {
		if int(Hero.Hero.Call("getBoundingClientRect").Get("left").Float()) > Arena.Perim.Ybottom {
			return nil
		}

		if _, ok := Path.Path[strconv.Itoa(Hero.PathCurrentIn.Num1)+strconv.Itoa(Hero.PathCurrentIn.Num2-1)]; !ok {
			return nil
		}
	}

	Hero.Position.Xleft += 10 * int(args[0].Float())

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.Xleft) + "px"))

	return nil
}

func MoveHeroY(this js.Value, args []js.Value) interface{} {

	if int(args[0].Float()) == 1 {
		if int(Hero.Hero.Call("getBoundingClientRect").Get("bottom").Float()) > Arena.Perim.Ybottom {
			return nil
		}

		if _, ok := Path.Path[strconv.Itoa(Hero.PathCurrentIn.Num1+1)+strconv.Itoa(Hero.PathCurrentIn.Num2)]; !ok {
			return nil
		}
	}
	if int(args[0].Float()) == -1 {
		if int(Hero.Hero.Call("getBoundingClientRect").Get("top").Float()) > Arena.Perim.Ybottom {
			return nil
		}

		if _, ok := Path.Path[strconv.Itoa(Hero.PathCurrentIn.Num1-1)+strconv.Itoa(Hero.PathCurrentIn.Num2)]; !ok {
			return nil
		}
	}

	Hero.Position.Ytop += 10 * int(args[0].Float())

	Hero.Style.Call("setProperty", "top", (strconv.Itoa(Hero.Position.Ytop) + "px"))

	return nil
}

func TakeDimensionScreen() {
	wind := js.Global().Get("window")

	Screen.Num1 = int(wind.Get("innerHeight").Float())
	Screen.Num2 = int(wind.Get("innerWidth").Float())

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
