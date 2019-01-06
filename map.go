package main

type Map struct {
	Width, Height int
	Cells         [][]*Cell
}

type Cell struct {
	Food int    `json:"food,omitempty"`
	Hive string `json:"hive,omitempty"`
	Ant  string `json:"ant,omitempty"`
}

type Object struct {
	x, y, food int
	used, hive bool
}

func (f *Object) distance(y, x int) int {
	w, h := 0, 0

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

func (m *Map) getObjects(id string) []Object {
	all := make([]Object, 0)
	for y, row := range m.Cells {
		for x, c := range row {
			if c.Hive == id {
				all = append(all, Object{y: y, x: x, hive: true})
				continue
			}
			if c.Food > 0 {
				all = append(all, Object{y: y, x: x, food: c.Food})
			}
		}
	}
	return all
}
