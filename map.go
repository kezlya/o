package main

type Map struct {
	Width, Height int
	Cells         [][]*Cell
	objects       []*Object
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

func (m *Map) getObjects(id string) {
	all := make([]*Object, 0)
	for y, row := range m.Cells {
		for x, c := range row {
			if c.Hive == id {
				all = append(all, &Object{y: y, x: x, hive: true})
				continue
			}
			if c.Food > 0 && (c.Hive == "" || c.Hive == id) {
				all = append(all, &Object{y: y, x: x, food: c.Food})
			}
		}
	}
	m.objects = all
}

func (m *Map) isEatable(y, x int, id string) bool {
	return y >= 0 && x >= 0 &&
		y < m.Height && x < m.Width &&
		m.Cells[y][x].Food > 0 &&
		(m.Cells[y][x].Hive == "" || m.Cells[y][x].Hive == id)
}

func (m *Map) isEmpty(y, x int, id string) bool {
	return m.Cells[y][x].Food == 0 && m.Cells[y][x].Ant == "" &&
		(m.Cells[y][x].Hive == "" || m.Cells[y][x].Hive == id)
}
