package main

func main() {
	StartServer()
}

func whatToDo(hive *Hive) map[int]BotOder {

	actions := make(map[int]BotOder)
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
		a.order = &BotOder{Unload, Up}
		return true
	}

	if a.Payload > 0 && a.X < a.hive.Map.Width-1 &&
		a.hive.Map.Cells[a.Y][a.X+1].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y][a.X+1].Ant == "" {
		a.order = &BotOder{Unload, Right}
		return true
	}

	if a.Payload > 0 && a.Y < a.hive.Map.Height-1 &&
		a.hive.Map.Cells[a.Y+1][a.X].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y+1][a.X].Ant == "" {
		a.order = &BotOder{Unload, Down}
		return true
	}

	if a.Payload > 0 && a.X > 0 &&
		a.hive.Map.Cells[a.Y][a.X-1].Hive == a.hive.Id &&
		a.hive.Map.Cells[a.Y][a.X-1].Ant == "" {
		a.order = &BotOder{Unload, Left}
		return true
	}

	return false
}

func (a *Ant) consume() bool {
	order := BotOder{}

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
		a.hive.Map.Cells[a.Y-1][a.X].Hive == "" {
		order.Dir = Right
		a.order = &order
		return true
	}

	if a.Y < a.hive.Map.Height-1 &&
		a.hive.Map.Cells[a.Y+1][a.X].Food > 0 &&
		a.hive.Map.Cells[a.Y-1][a.X].Hive == "" {
		order.Dir = Down
		a.order = &order
		return true
	}

	if a.X > 0 &&
		a.hive.Map.Cells[a.Y][a.X-1].Food > 0 &&
		a.hive.Map.Cells[a.Y-1][a.X].Hive == "" {
		order.Dir = Left
		a.order = &order
		return true
	}

	return false
}

func (a *Ant) move() {
	zoom := 1
	ring := a.hive.Map.around(a.Y, a.X, zoom)

	for len(ring) > 0 && a.order == nil {
		for _, cell := range ring {
			if a.Payload == 9 { // go home
				if cell.Hive == a.hive.Id {
					if a.direction(cell.y, cell.x) {
						break
					}
				}
			} else if a.Payload > 5 { // whatever first home or food
				if cell.Hive == a.hive.Id || cell.Food > 0 {
					if a.direction(cell.y, cell.x) {
						break
					}
				}
			} else { // search for food
				if cell.Hive == "" && cell.Food > 0 {
					if a.direction(cell.y, cell.x) {
						break
					}
				}
			}
		}
		zoom++
		ring = a.hive.Map.around(a.Y, a.X, zoom)
	}
}

func (m *Map) around(oy, ox uint, zoom int) []*Cell {
	ring := make([]*Cell, 0)
	for y := -zoom; y <= zoom; y++ {
		for x := -zoom; x <= zoom; x++ {
			if y == 0 && x == 0 {
				continue
			}
			c := m.isValid(uint(y)+oy, uint(x)+ox)
			if c != nil {
				ring = append(ring, c)
			}
		}
	}
	return ring
}

func (m *Map) isValid(y, x uint) *Cell {
	if y < m.Height && x < m.Width {
		m.Cells[y][x].y = y
		m.Cells[y][x].x = x
		return m.Cells[y][x]
	}
	return nil
}

//TODO: check for future occupied cells by my ants
func (a *Ant) direction(dy, dx uint) bool {
	if a.X < dx && a.canMove(a.Y, a.X+1) {
		a.order = &BotOder{Move, Right}
		return true
	}

	if a.Y < dy && a.canMove(a.Y+1, a.X) {
		a.order = &BotOder{Move, Down}
		return true
	}

	if a.X > dx && a.canMove(a.Y, a.X-1) {
		a.order = &BotOder{Move, Left}
		return true
	}

	if a.Y > dy && a.canMove(a.Y-1, a.X) {
		a.order = &BotOder{Move, Up}
		return true
	}
	return false
}

func (a *Ant) canMove(dy, dx uint) bool {
	return (a.hive.Map.Cells[dy][dx].Hive == "" || a.hive.Map.Cells[dy][dx].Hive == a.hive.Id) &&
		a.hive.Map.Cells[dy][dx].Food == 0 && a.hive.Map.Cells[dy][dx].Ant == ""
}
