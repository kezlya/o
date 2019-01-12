package main

type Hive struct {
	Tick int
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
