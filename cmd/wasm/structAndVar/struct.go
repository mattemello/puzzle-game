package structandvar

import "syscall/js"

type Coordination struct {
	Xleft   int
	Xright  int
	Ytop    int
	Ybottom int
}

type Screen struct {
	Width  int
	Height int
}

type TheArena struct {
	DimensionCol int
	DimensionRaw int
	Perim        Coordination
}

type TheHero struct {
	Hero     js.Value
	Ctxh     js.Value
	Style    js.Value
	Position Coordination
	//Dimension Coordination
}
