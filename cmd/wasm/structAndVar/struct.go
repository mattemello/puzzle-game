package structandvar

import "syscall/js"

type Coordination struct {
	Xleft   int
	Xright  int
	Ytop    int
	Ybottom int
}

type Path struct {
	Arena        js.Value
	Ctx          js.Value
	Numer1       int
	Numer2       int
	Coordination Coordination
}

type ControllPath struct {
	Path      map[string]Path
	ArrayPath []string
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
