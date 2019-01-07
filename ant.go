package main

type Ant struct {
	Wasted, Age, Health int
	Payload, X, Y       int
	Event               string

	hive  *Hive
	order *AntOder
}

type AntOder struct {
	Act
	Dir
}

type Act string

const (
	Stay   Act = "stay"
	Move   Act = "move"
	Load   Act = "load"
	Unload Act = "unload"
	Eat    Act = "eat"
)

type Dir string

const (
	Up    Dir = "up"
	Right Dir = "right"
	Down  Dir = "down"
	Left  Dir = "left"
)

//TODO: check for future occupied cells by my ants
func (a *Ant) direction(dy, dx int) bool {
	if a.X < dx && a.canMove(a.Y, a.X+1) {
		a.order = &AntOder{Move, Right}
		return true
	}

	if a.Y < dy && a.canMove(a.Y+1, a.X) {
		a.order = &AntOder{Move, Down}
		return true
	}

	if a.X > dx && a.canMove(a.Y, a.X-1) {
		a.order = &AntOder{Move, Left}
		return true
	}

	if a.Y > dy && a.canMove(a.Y-1, a.X) {
		a.order = &AntOder{Move, Up}
		return true
	}
	return false
}

func (a *Ant) canMove(dy, dx int) bool {
	return (a.hive.Map.Cells[dy][dx].Hive == "" || a.hive.Map.Cells[dy][dx].Hive == a.hive.Id) &&
		a.hive.Map.Cells[dy][dx].Food == 0 && a.hive.Map.Cells[dy][dx].Ant == ""
}
