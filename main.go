package main

type Hive struct {
	Id   string
	Ants map[int]*Ant
	Map  *Map
}

func main() {
	StartServer()
}

func whatToDo(hive *Hive) map[int]AntOder {

	actions := make(map[int]AntOder)
	for id, ant := range hive.Ants {
		ant.hive = hive

		if ant.unload() {
			actions[id] = *ant.order
			continue
		}

		if ant.consume() {
			actions[id] = *ant.order
			continue
		}

		ant.move()
		if ant.order != nil {
			actions[id] = *ant.order
		}
	}

	return actions
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

	if a.Y > 0 &&
		a.hive.Map.Cells[a.Y-1][a.X].Food > 0 &&
		a.hive.Map.Cells[a.Y-1][a.X].Hive == "" {
		order.Dir = Up
		a.order = &order
		return true
	}

	if a.X < a.hive.Map.Width-1 &&
		a.hive.Map.Cells[a.Y][a.X+1].Food > 0 &&
		a.hive.Map.Cells[a.Y][a.X+1].Hive == "" {
		order.Dir = Right
		a.order = &order
		return true
	}

	if a.Y < a.hive.Map.Height-1 &&
		a.hive.Map.Cells[a.Y+1][a.X].Food > 0 &&
		a.hive.Map.Cells[a.Y+1][a.X].Hive == "" {
		order.Dir = Down
		a.order = &order
		return true
	}

	if a.X > 0 &&
		a.hive.Map.Cells[a.Y][a.X-1].Food > 0 &&
		a.hive.Map.Cells[a.Y][a.X-1].Hive == "" {
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
				firstTarget = object
			} else {
				secondTarget = object
			}
		}
	}

	if firstTarget != nil && a.direction(firstTarget.y, firstTarget.x) {
		return
	}

	if secondTarget != nil {
		a.direction(secondTarget.y, secondTarget.x)
	}
}
