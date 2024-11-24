package createarena

import (
	strctVar "github.com/mattemello/puzzle-game/cmd/wasm/structAndVar"
	"strconv"
	"syscall/js"
)

var Screen strctVar.Screen
var Arena strctVar.TheArena
var doc = js.Global().Get("document")

// TODO: scale with the monitor resolution;

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

func TakeDimensionScreen() {
	wind := js.Global().Get("window")

	Screen.Num1 = int(wind.Get("innerHeight").Float())
	Screen.Num2 = int(wind.Get("innerWidth").Float())

}
