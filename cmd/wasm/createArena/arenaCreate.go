package createarena

import (
	"syscall/js"

	strctVar "github.com/mattemello/puzzle-game/cmd/wasm/structAndVar"
)

var Screen strctVar.Screen
var Arena strctVar.TheArena
var doc = js.Global().Get("document")

// TODO: scale with the monitor resolution;

func CreateTheArena(this js.Value, args []js.Value) interface{} {

	Arena.DimensionCol, Arena.DimensionRaw = TakeDimensionScreen()

	js.Global().Call("Arena", Arena.DimensionRaw, Arena.DimensionCol)

	TakedimensionArena()
	CreateThePath()

	return nil
}

func TakedimensionArena() {

	perimeterAre := doc.Call("getElementById", "arena0-0")
	Arena.Perim.Xleft = int(perimeterAre.Call("getBoundingClientRect").Get("left").Float())
	Arena.Perim.Ytop = int(perimeterAre.Call("getBoundingClientRect").Get("top").Float())

	perimeterAre = doc.Call("getElementById", "arena"+CalculateKey(0, Arena.DimensionCol-1))
	Arena.Perim.Xright = int(perimeterAre.Call("getBoundingClientRect").Get("right").Float())

	perimeterAre = doc.Call("getElementById", "arena"+CalculateKey(Arena.DimensionRaw-1, Arena.DimensionCol-1))
	Arena.Perim.Ybottom = int(perimeterAre.Call("getBoundingClientRect").Get("bottom").Float())

}

func TakeDimensionScreen() (int, int) {
	wind := js.Global().Get("window")

	Screen.Num1 = int(wind.Get("innerHeight").Float()) / 50
	Screen.Num2 = int(wind.Get("innerWidth").Float()) / 50

	if Screen.Num1 < Screen.Num2 {
		return Screen.Num1, Screen.Num1
	}

	return Screen.Num2, Screen.Num2

}
