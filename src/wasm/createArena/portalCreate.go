package createarena

import (
	"math/rand"
	"syscall/js"
)

//todo: maybe create the portal in the more distant place from the start

func createPortaleWin() {
	value := 0

	for Path.ArrayPath[value] == Path.ArrayPath[0] {
		value = rand.Intn(len(Path.ArrayPath))
	}

	Path.Path[Path.ArrayPath[value]].Portal = true

}

func OnThePortal() {
	js.Global().Call("levelCompleate")
}
