package createarena

import (
	"syscall/js"
)

func assoluteValue(num int) int {
	if num < 0 {
		return (num * -1)
	}

	return num
}

func createPortaleWin() {

	maxDistance := -1
	saveKey := ""

	for key, elem := range Path.Path {

		num1 := assoluteValue(Path.Path[Path.ArrayPath[0]].Number1 - elem.Number1)
		num2 := assoluteValue(Path.Path[Path.ArrayPath[0]].Number2 - elem.Number2)

		if (num1 + num2) > maxDistance {
			saveKey = key
			maxDistance = (num1 + num2)
		}

	}

	Path.Path[saveKey].Portal = true

}

func OnThePortal() {
	js.Global().Call("levelCompleate")
}
