package structandvar

import "syscall/js"

type Coordination struct {
	X int
	Y int
}

type TheArena struct {
	Dimension Coordination
}

type TheHero struct {
	Hero     js.Value
	Ctxh     js.Value
	Style    js.Value
	Position Coordination
	//Dimension Coordination
}
