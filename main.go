package main

import (
	"math/rand"
	"time"
)

type BotOder struct {
	Act    string
	Dir string
}

func main() {
	StartServer()
}

func whatToDo(hive *Hive) map[int]BotOder {
	actions := make(map[int]BotOder)
	rand.Seed(time.Now().UnixNano())
	for id, ant := range hive.Ants {

		//Default action if ant don't see Food
		action := "move"

		//Default direction is Random
		direction := []string{"up","down","left","right"}[rand.Intn(4)]
		food, hive, dir  := lookAround(ant, hive.Map)

		if hive && ant.Payload>0{
			direction = dir
			action = "unload"
		}else if food{
			direction = dir
			if ant.Health<9{action = "eat"}
			if ant.Payload<9 {action = "load"}
		}

		actions[id] = BotOder{action, direction}
		//time.Sleep(400*time.Millisecond)
	}

	return actions
}

func lookAround(ant *Ant, world *Map)(food, hive bool, dir string){

	if ant.Y > 0 {
		dir = "u"
		food,hive = iSee(ant.Y-1,ant.X,world)
	}
	
	if ant.Y < world.Height-1{
		dir = "d"
		food,hive = iSee(ant.Y+1,ant.X,world)
	}

	if ant.X < world.Width-1{
		dir = "r"
		food,hive = iSee(ant.Y,ant.X+1,world)
	}

	if ant.X >0{
		dir = "l"
		food,hive = iSee(ant.Y,ant.X-1,world)
	}

	return
}

func iSee(y,x uint8, world *Map) (food, hive bool) {
	if world.Cells[y][x].Food>0{
		food = true
	}
	if world.Cells[y][x].CellType > 0 && world.Cells[y][x].CellType <7{
		hive = true
	}
	return
}