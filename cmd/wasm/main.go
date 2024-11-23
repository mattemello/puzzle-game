package main

import (
	//"fmt"
	"fmt"
	"strconv"
	"syscall/js"

	pathArena "github.com/mattemello/puzzle-game/cmd/wasm/createArena"
	strctVar "github.com/mattemello/puzzle-game/cmd/wasm/structAndVar"
)

/* NOTE: the approssimative position of the circle hero:
 * right: ~(-40)
 * left: ~(-30)
 * top: ~(-8)
 * bottom: ~(-8)
 * it takes 4 step to pass a block, so if we want to change block we need to do like at 2, with in memory the prec one, and controll the border evrytime
 * i have to implement a perimeter of the path in the hero maybe
 */

/**
 * TODO: make the path black, so that it will create when you walk it
 * maybe its possible to take the element near the hero, save them and say if they are walkable or not
 * so i can color the element walkable
 **/

var Hero strctVar.TheHero
var move strctVar.Coordination

var doc = js.Global().Get("document")

func CreateTheHero(this js.Value, args []js.Value) interface{} {

	Hero.Position.Xleft = pathArena.Path.Path[pathArena.Path.ArrayPath[0]].Coordination.Xleft - 40
	Hero.Position.Ytop = pathArena.Path.Path[pathArena.Path.ArrayPath[0]].Coordination.Ytop

	Hero.PathCurrentIn.Num1 = pathArena.Path.Path[pathArena.Path.ArrayPath[0]].Number1
	Hero.PathCurrentIn.Num2 = pathArena.Path.Path[pathArena.Path.ArrayPath[0]].Number2

	Hero.Style = Hero.Hero.Get("style")

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.Xleft) + "px"))
	Hero.Style.Call("setProperty", "top", (strconv.Itoa(Hero.Position.Ytop) + "px"))
	js.Global().Call("Hero", Hero.Ctxh)

	return nil
}

func MoveHeroX(this js.Value, args []js.Value) interface{} {

	if int(args[0].Float()) == 1 {
		if int(Hero.Hero.Call("getBoundingClientRect").Get("right").Float())-20 > pathArena.Arena.Perim.Xright {
			return nil
		}

		heroSurpassTheBlock := int(Hero.Hero.Call("getBoundingClientRect").Get("right").Float())-20 > pathArena.Path.Path[pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2)].Coordination.Xright
		if heroSurpassTheBlock {
			if !pathArena.Path.Path[pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2)].Wall.Xright {
				return nil
			}

			Hero.PathCurrentIn.Num2 = Hero.PathCurrentIn.Num2 + 1
		}

	}
	if int(args[0].Float()) == -1 {
		if int(Hero.Hero.Call("getBoundingClientRect").Get("left").Float())+30 < pathArena.Arena.Perim.Xleft {
			return nil
		}

		fmt.Println(pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2))
		heroSurpassTheBlock := int(Hero.Hero.Call("getBoundingClientRect").Get("left").Float())-30 < pathArena.Path.Path[pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2)].Coordination.Xleft
		if heroSurpassTheBlock {
			if !pathArena.Path.Path[pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2)].Wall.Xleft {
				//fmt.Println(int(Hero.Hero.Call("getBoundingClientRect").Get("left").Float())-30, pathArena.Path.Path[pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2)].Coordination.Xleft)
				return nil
			}

			Hero.PathCurrentIn.Num2 = Hero.PathCurrentIn.Num2 - 1
		}

	}

	Hero.Position.Xleft += 10 * int(args[0].Float())

	Hero.Style.Call("setProperty", "left", (strconv.Itoa(Hero.Position.Xleft) + "px"))

	return nil
}

func MoveHeroY(this js.Value, args []js.Value) interface{} {

	if int(args[0].Float()) == 1 {
		if int(Hero.Hero.Call("getBoundingClientRect").Get("bottom").Float()) > pathArena.Arena.Perim.Ybottom {
			return nil
		}

		//fmt.Println(pathArena.Path.Path[pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2)].Coordination.Ybottom)
		heroSurpassTheBlock := int(Hero.Hero.Call("getBoundingClientRect").Get("bottom").Float()) > pathArena.Path.Path[pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2)].Coordination.Ybottom
		if heroSurpassTheBlock {
			if !pathArena.Path.Path[pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2)].Wall.Ybottom {
				return nil
			}

			Hero.PathCurrentIn.Num1 = Hero.PathCurrentIn.Num1 + 1
		}
	}
	if int(args[0].Float()) == -1 {
		if int(Hero.Hero.Call("getBoundingClientRect").Get("top").Float()) < pathArena.Arena.Perim.Ytop {
			return nil
		}

		heroSurpassTheBlock := int(Hero.Hero.Call("getBoundingClientRect").Get("top").Float()) < pathArena.Path.Path[pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2)].Coordination.Ytop
		if heroSurpassTheBlock {
			if !pathArena.Path.Path[pathArena.CalculateKey(Hero.PathCurrentIn.Num1, Hero.PathCurrentIn.Num2)].Wall.Ytop {
				return nil
			}

			Hero.PathCurrentIn.Num1 = Hero.PathCurrentIn.Num1 - 1
		}
	}

	Hero.Position.Ytop += 10 * int(args[0].Float())

	Hero.Style.Call("setProperty", "top", (strconv.Itoa(Hero.Position.Ytop) + "px"))

	return nil
}

func main() {

	pathArena.TakeDimensionScreen()

	Hero.Hero = doc.Call("getElementById", "hero")
	Hero.Ctxh = Hero.Hero.Call("getContext", "2d")

	js.Global().Set("CreateTheArena", js.FuncOf(pathArena.CreateTheArena))
	js.Global().Set("CreateTheHero", js.FuncOf(CreateTheHero))
	js.Global().Set("MoveHeroX", js.FuncOf(MoveHeroX))
	js.Global().Set("MoveHeroY", js.FuncOf(MoveHeroY))

	<-make(chan struct{})
}
