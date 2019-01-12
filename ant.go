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
	if a.X < dx && a.hive.Map.isEmpty(a.Y, a.X+1, a.hive.Id) {
		a.order = &AntOder{Move, Right}
		return true
	}

	if a.Y < dy && a.hive.Map.isEmpty(a.Y+1, a.X, a.hive.Id) {
		a.order = &AntOder{Move, Down}
		return true
	}

	if a.X > dx && a.hive.Map.isEmpty(a.Y, a.X-1, a.hive.Id) {
		a.order = &AntOder{Move, Left}
		return true
	}

	if a.Y > dy && a.hive.Map.isEmpty(a.Y-1, a.X, a.hive.Id) {
		a.order = &AntOder{Move, Up}
		return true
	}
	return false
}

func (a *Ant) unload() bool {
	if a.Payload > 0 && a.Y > 0 &&
		a.hive.Map.Cells[a.Y-1][a.X].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y-1][a.X].Ant == "" {
		a.order = &AntOder{Unload, Up}
		return true
	}

	if a.Payload > 0 && a.X < a.hive.Map.Width-1 &&
		a.hive.Map.Cells[a.Y][a.X+1].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y][a.X+1].Ant == "" {
		a.order = &AntOder{Unload, Right}
		return true
	}

	if a.Payload > 0 && a.Y < a.hive.Map.Height-1 &&
		a.hive.Map.Cells[a.Y+1][a.X].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y+1][a.X].Ant == "" {
		a.order = &AntOder{Unload, Down}
		return true
	}

	if a.Payload > 0 && a.X > 0 &&
		a.hive.Map.Cells[a.Y][a.X-1].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y][a.X-1].Ant == "" {
		a.order = &AntOder{Unload, Left}
		return true
	}

	return false
}

func (a *Ant) consume() bool {
	order := AntOder{}

	if a.Health < 9 {
		order.Act = Eat
	} else if a.Payload < 9 {
		order.Act = Load
	} else {
		return false
	}

	if a.hive.Map.isEatable(a.Y-1, a.X, a.hive.Id) {
		order.Dir = Up
		a.order = &order
		return true
	}

	if a.hive.Map.isEatable(a.Y, a.X+1, a.hive.Id) {
		order.Dir = Right
		a.order = &order
		return true
	}

	if a.hive.Map.isEatable(a.Y+1, a.X, a.hive.Id) {
		order.Dir = Down
		a.order = &order
		return true
	}

	if a.hive.Map.isEatable(a.Y, a.X-1, a.hive.Id) {
		order.Dir = Left
		a.order = &order
		return true
	}

	return false
}

func (a *Ant) move() {

	shortest := 9999999
	var firstTarget *Object
	var secondTarget *Object
	for _, object := range a.hive.Map.objects {
		if a.Payload == 9 && !object.hive { // move home
			continue
		}

		if a.Payload < 5 && object.hive { // search for food
			continue
		}

		s := object.distance(a.Y, a.X)
		if s == 0 {
			continue
		}

		if s < shortest {
			shortest = s
			if !object.used {
				secondTarget = firstTarget
				firstTarget = object
			} else {
				secondTarget = object
			}
		}
	}

	if firstTarget != nil &&
		a.direction(firstTarget.y, firstTarget.x) {
		return
	}

	if secondTarget != nil &&
		a.direction(secondTarget.y, secondTarget.x) {
		return
	}

	a.order = &AntOder{Act: Stay}
}
