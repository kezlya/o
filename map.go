package main

import "github.com/kezlya/anthive"

type Object struct {
	x, y, food uint
	used, hive bool
}

func (f *Object) distance(y, x uint) uint {
	var w, h uint

	if f.x > x {
		w = f.x - x
	} else {
		w = x - f.x
	}

	if f.y > y {
		h = f.y - y
	} else {
		h = y - f.y
	}

	return w + h
}

func isFood(y, x uint, order *anthive.Order) bool {
	if y >= 0 && x >= 0 &&
		y < canvas.Height && x < canvas.Width &&
		canvas.Cells[y][x].Food > 0 {
		if order.Action == anthive.ActionLoad && canvas.Cells[y][x].Hive != "" {
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

func getObjects() []*Object {
	all := make([]*Object, 0)
	for y, row := range canvas.Cells {
		for x, c := range row {
			if c.Hive == id {
				all = append(all, &Object{y: uint(y), x: uint(x), hive: true})
				continue
			}
			if c.Food > 0 && (c.Hive == "" || c.Hive == id) {
				all = append(all, &Object{y: uint(y), x: uint(x), food: c.Food})
			}
		}
	}
	return all
}
