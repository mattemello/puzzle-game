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

/* type MarginPath struct {
} */

type ControllPath struct {
	Path      map[string]Path
	ArrayPath []string
}

type Screen struct {
	Num1 int
	Num2 int
}

type TheArena struct {
	DimensionCol int
	DimensionRaw int
	Perim        Coordination
}

type TheHero struct {
	Hero          js.Value
	Ctxh          js.Value
	Style         js.Value
	Position      Coordination
	PathCurrentIn Screen
	//Dimension Coordination
}
