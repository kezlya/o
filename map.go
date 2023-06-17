package main

type Target struct {
	x, y, food uint
	used, hive bool
}

func (t *Target) distance(y, x uint) uint {
	var w, h uint

	if t.x > x {
		w = t.x - x
	} else {
		w = x - t.x
	}

	if t.y > y {
		h = t.y - y
	} else {
		h = y - t.y
	}

	return w + h
}

func isFood(y, x uint, order *Order) bool {
	if y >= 0 && x >= 0 &&
		y < canvas.Height && x < canvas.Width &&
		canvas.Cells[y][x].Food > 0 {
		if order.Action == ActionTake && canvas.Cells[y][x].Hive != "" {
			return false
		}
		return true
	}
	return false
}

func isEmpty(y, x uint) bool {
	return canvas.Cells[y][x].Food == 0 && canvas.Cells[y][x].Ant == "" &&
		(canvas.Cells[y][x].Hive == "" || canvas.Cells[y][x].Hive == id)
}

func getTargets() []*Target {
	all := make([]*Target, 0)
	for y, row := range canvas.Cells {
		for x, c := range row {
			if c.Hive == id {
				all = append(all, &Target{y: uint(y), x: uint(x), hive: true})
				continue
			}
			if c.Food > 0 && (c.Hive == "" || c.Hive == id) {
				all = append(all, &Target{y: uint(y), x: uint(x), food: c.Food})
			}
		}
	}
	return all
}
