package main

func main() {
	StartServer() //Breaking code
}

func whatToDo(hive *Hive) map[int]BotOder {
	//food := hive.Map.getFood()
	//home := hive.Map.getHome(hive.Id)

	actions := make(map[int]BotOder)

	for id, ant := range hive.Ants {
		//antPoint := point{y: ant.Y, x: ant.X}

		if ant.unload() {
			actions[id] = *ant.order
			continue
		}

		if ant.consume() {
			actions[id] = *ant.order
			continue
		}

		//if ant.Payload < 9 {
		//	actions[id] = BotOder{Move, antPoint.towardsFood(food, hive.Map)}
		//	continue
		//}
		//
		//actions[id] = BotOder{Move, antPoint.towardsHome(home, hive.Map)}
	}

	return actions
}

func (a *Ant) unload() bool {
	if a.hive.Map.Cells[a.Y-1][a.X].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y-1][a.X].Ant == "" {
		a.order = &BotOder{Unload, Up}
		return true
	}

	if a.X < a.hive.Map.Width-1 &&
		a.hive.Map.Cells[a.Y][a.X+1].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y][a.X+1].Ant == "" {
		a.order = &BotOder{Unload, Right}
		return true
	}

	if a.Y < a.hive.Map.Height-1 &&
		a.hive.Map.Cells[a.Y+1][a.X].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y+1][a.X].Ant == "" {
		a.order = &BotOder{Unload, Down}
		return true
	}

	if a.hive.Map.Cells[a.Y][a.X-1].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y][a.X-1].Ant == "" {
		a.order = &BotOder{Unload, Left}
		return true
	}

	return false
}

func (a *Ant) consume() bool {
	if a.hive.Map.Cells[a.Y-1][a.X].Food > 0 {
		a.order = &BotOder{Load, Up}
	}

	if a.X < a.hive.Map.Width-1 &&
		a.hive.Map.Cells[a.Y][a.X+1].Food > 0 {
		a.order = &BotOder{Load, Right}
	}

	if a.Y < a.hive.Map.Height-1 &&
		a.hive.Map.Cells[a.Y+1][a.X].Food > 0 {
		a.order = &BotOder{Load, Down}
	}

	if a.hive.Map.Cells[a.Y][a.X-1].Food > 0 {
		a.order = &BotOder{Load, Left}
	}

	if a.order != nil {
		if a.Health < 9 {
			// check that they don't eat from home
			a.order.Act = Eat
		}
		return true
	}

	return false
}

func (a *Ant) around(oy, ox, zoom uint) []*Cell {
	ring := make([]*Cell, 0)
	for y := -zoom; y <= zoom; y++ {
		for x := -zoom; x <= zoom; x++ {
			if y == 0 && x == 0 {
				continue
			}
			c := a.hive.getCell(y+oy, x+ox)
			if c != nil {
				ring = append(ring, c)
			}
		}
	}
	return ring
}

func (h *Hive) getCell(y, x uint) *Cell {
	if y < h.Map.Height && x < h.Map.Width {
		return h.Map.Cells[y][x]
	}
	return nil
}

func (m *Map) getFood() map[point]int {
	food := make(map[point]int)
	for y, row := range m.Cells {
		for x, c := range row {
			if c.Food > 0 {
				food[point{y: y, x: x}] = c.Food
			}
		}
	}
	return food
}

func (m *Map) getHome(id string) []point {
	home := make([]point, 0)
	for y, row := range m.Cells {
		for x, c := range row {
			if c.Hive == id {
				home = append(home, point{y: y, x: x})
			}
		}
	}
	return home
}

func (m *Map) isCellEmpty(p point) bool {
	return m.Cells[p.y][p.x].Food == 0 &&
		m.Cells[p.y][p.x].Ant == "" &&
		m.Cells[p.y][p.x].Hive == ""
}

func (p *point) towardsFood(f map[point]int, world *Map) Dir {
	effort := 10000000
	dir := Up
	for to := range f {
		ticks := p.distance(to)
		if ticks < effort {
			dir = p.move(to, world)
			effort = ticks
		}
	}
	return dir
}

func (p *point) towardsHome(h []point, world *Map) Dir {
	effort := 10000000
	dir := Up
	for _, to := range h {
		ticks := p.distance(to)
		if ticks < effort {
			dir = p.move(to, world)
			effort = ticks
		}
	}
	return dir
}

func (from *point) distance(to point) int {
	w, h := 0, 0
	if from.x > to.x {
		w = from.x - to.x
	} else {
		w = to.x - from.x
	}
	if from.y > to.y {
		h = from.y - to.y
	} else {
		h = to.y - from.y
	}
	return w + h
}

func (from *point) move(to point, world *Map) Dir {
	if from.x > to.x && world.isCellEmpty(point{x: from.x - 1, y: from.y}) {
		return Left
	}
	if from.x < to.x && world.isCellEmpty(point{x: from.x + 1, y: from.y}) {
		return Right
	}
	if from.y > to.y && world.isCellEmpty(point{x: from.x, y: from.y - 1}) {
		return Up
	}
	// from.y < to.y
	return Down
}
