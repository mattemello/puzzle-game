package createarena

import (
	"math/rand"
	"strconv"
	"syscall/js"

	error "github.com/mattemello/puzzle-game/cmd/wasm/Errors"
	strctVar "github.com/mattemello/puzzle-game/cmd/wasm/structAndVar"
)

var Path strctVar.ControllPath

func CalculateKey(num1, num2 int) string {
	return strconv.Itoa(num1) + "-" + strconv.Itoa(num2)
}

func ColorWhenPass(num1, num2 int) {

	if !Path.Path[CalculateKey(num1, num2)].Wall.Ybottom {
		js.Global().Call("colorPath", Path.Path[CalculateKey(num1+1, num2)].Ctx, "#313244")
	}
	if !Path.Path[CalculateKey(num1, num2)].Wall.Ytop {
		js.Global().Call("colorPath", Path.Path[CalculateKey(num1-1, num2)].Ctx, "#313244")
	}
	if !Path.Path[CalculateKey(num1, num2)].Wall.Xright {
		js.Global().Call("colorPath", Path.Path[CalculateKey(num1, num2+1)].Ctx, "#313244")
	}
	if !Path.Path[CalculateKey(num1, num2)].Wall.Xleft {
		js.Global().Call("colorPath", Path.Path[CalculateKey(num1, num2-1)].Ctx, "#313244")
	}

}

// TODO: make not do duplicate

func chooseThePath(block strctVar.Path, dimensioPathNow int) strctVar.Path {

	possibleDirection := make(map[int]string)

	possibleDirection[0] = "up"
	possibleDirection[1] = "down"
	possibleDirection[2] = "left"
	possibleDirection[3] = "right"

	// NOTE: number2 -> left/right number1 -> up/down
	if block.Number1 == 0 {
		delete(possibleDirection, 0)
	} else if block.Number1 == Arena.DimensionCol-1 {
		delete(possibleDirection, 1)
	}

	if block.Number2 == 0 {
		delete(possibleDirection, 2)
	} else if block.Number2 == Arena.DimensionRaw-1 {
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

	switch decision {

	case "up":
		newBlock.Number1 = block.Number1 - 1
		newBlock.Number2 = block.Number2
		break
	case "down":
		newBlock.Number1 = block.Number1 + 1
		newBlock.Number2 = block.Number2

		break
	case "left":
		newBlock.Number2 = block.Number2 - 1
		newBlock.Number1 = block.Number1

		break
	case "right":
		newBlock.Number2 = block.Number2 + 1
		newBlock.Number1 = block.Number1

		break

	default:
		js.Global().Call("alert", "Error - Error in the creation of the path, value not good")
		js.Global().Get("console").Call("error", "Error - Error in the creation of the path, value not good")
	}

	newBlock.Arena = doc.Call("getElementById", "arena"+CalculateKey(newBlock.Number1, newBlock.Number2))

	newBlock.Coordination.Xleft = int(newBlock.Arena.Call("getBoundingClientRect").Get("left").Float())
	newBlock.Coordination.Xright = int(newBlock.Arena.Call("getBoundingClientRect").Get("right").Float())
	newBlock.Coordination.Ybottom = int(newBlock.Arena.Call("getBoundingClientRect").Get("bottom").Float())
	newBlock.Coordination.Ytop = int(newBlock.Arena.Call("getBoundingClientRect").Get("top").Float())

	newBlock.Wall.Xleft = true
	newBlock.Wall.Ytop = true
	newBlock.Wall.Xright = true
	newBlock.Wall.Ybottom = true

	newBlock.Ctx = newBlock.Arena.Call("getContext", "2d")
	js.Global().Call("colorPath", newBlock.Ctx, "#313244")

	return newBlock
}

func CreateThePath() {

	dimensionPath := 0

	for !(dimensionPath > (Arena.DimensionCol) && dimensionPath < (Arena.DimensionCol*Arena.DimensionRaw)/3) {
		dimensionPath = rand.Intn(Arena.DimensionRaw * Arena.DimensionCol)
	}

	Path.ArrayPath = make([]string, dimensionPath)
	Path.Path = make(map[string]*strctVar.Path)

	num1 := rand.Intn(Arena.DimensionCol)
	num2 := rand.Intn(Arena.DimensionRaw)

	aren := doc.Call("getElementById", "arena"+CalculateKey(num1, num2))

	var coo strctVar.Coordination
	coo.Xleft = int(aren.Call("getBoundingClientRect").Get("left").Float())
	coo.Xright = int(aren.Call("getBoundingClientRect").Get("right").Float())
	coo.Ybottom = int(aren.Call("getBoundingClientRect").Get("bottom").Float())
	coo.Ytop = int(aren.Call("getBoundingClientRect").Get("top").Float())

	ctx := aren.Call("getContext", "2d")
	js.Global().Call("colorPath", ctx, "#313244")

	Path.ArrayPath[0] = CalculateKey(num1, num2)

	Path.Path[Path.ArrayPath[0]] = &strctVar.Path{Number1: num1, Number2: num2, Arena: aren, Ctx: ctx, Coordination: coo, Wall: strctVar.Wall{Xleft: true, Xright: true, Ytop: true, Ybottom: true}}

	for i := 1; i < dimensionPath; i++ {
		block := chooseThePath(*Path.Path[Path.ArrayPath[i-1]], i)
		Path.ArrayPath[i] = CalculateKey(block.Number1, block.Number2)
		Path.Path[Path.ArrayPath[i]] = &block

	}

	var wall int
	for key, ele := range Path.Path {
		wall = 0

		if _, exist := Path.Path[CalculateKey(ele.Number1, ele.Number2+1)]; exist {
			Path.Path[key].Wall.Xright = false
			wall++
		}
		if _, exist := Path.Path[CalculateKey(ele.Number1, ele.Number2-1)]; exist {
			Path.Path[key].Wall.Xleft = false
			wall++
		}
		if _, exist := Path.Path[CalculateKey(ele.Number1+1, ele.Number2)]; exist {
			Path.Path[key].Wall.Ybottom = false
			wall++
		}
		if _, exist := Path.Path[CalculateKey(ele.Number1-1, ele.Number2)]; exist {
			Path.Path[key].Wall.Ytop = false
			wall++
		}

		error.AssertError(wall == 0, "error in the creation of the path")

	}

}
