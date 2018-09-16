package main

import (
	"math/rand"
	"time"
)

type BotOder struct {
	Act
	Dir
}

type Act string

const (
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

func main() {
	StartServer()
}

var username string

func whatToDo(hive *Hive) map[int]BotOder {
	username = hive.Username
	actions := make(map[int]BotOder)
	rand.Seed(time.Now().UnixNano())

	for id, ant := range hive.Ants {

		//Default action if ant don't see Food
		action := Move

		//Default direction is Random
		direction := []Dir{Up, Down, Left, Right}[rand.Intn(4)]
		food, hive, dir := lookAround(ant, hive.Map)

		if hive {
			if ant.Payload > 0 {
				direction = dir
				action = Unload
			}

		} else if food {
			direction = dir
			if ant.Health < 9 {
				action = Eat
			}
			if ant.Payload < 9 {
				action = Load
			}
		}

		actions[id] = BotOder{action, direction}

	}
	//time.Sleep(400 * time.Millisecond)

	return actions
}

func lookAround(ant *Ant, world *Map) (food, hive bool, dir Dir) {

	if ant.Y > 0 {
		dir = Up
		food, hive = iSee(world.Cells[ant.Y-1][ant.X])
	}

	if ant.Y < world.Height-1 {
		dir = Down
		food, hive = iSee(world.Cells[ant.Y+1][ant.X])
	}

	if ant.X < world.Width-1 {
		dir = Right
		food, hive = iSee(world.Cells[ant.Y][ant.X+1])
	}

	if ant.X > 0 {
		dir = Left
		food, hive = iSee(world.Cells[ant.Y][ant.X-1])
	}

	return
}

func iSee(cell *Cell) (food, hive bool) {
	if cell.Food > 0 {
		food = true
	}
	if cell.Hive == "username" {
		hive = true
	}
	return
}
