package structandvar

import "syscall/js"

type Coordination struct {
	Xleft   int
	Xright  int
	Ytop    int
	Ybottom int
}

type Wall struct {
	Xleft   bool
	Xright  bool
	Ytop    bool
	Ybottom bool
}

type Path struct {
	Arena        js.Value
	Ctx          js.Value
	Number1      int
	Number2      int
	Coordination Coordination
	Wall         Wall
	Portal       bool
}

/* type MarginPath struct {
} */

type ControllPath struct {
	Path      map[string]*Path
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
