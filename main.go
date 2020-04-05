package main

import "github.com/kezlya/anthive"

var id string
var canvas *anthive.Canvas

func main() {
	StartServer()
}

func whatToDo(request *anthive.BotRequest) map[uint16]*anthive.Order {
	orders := make(map[uint16]*anthive.Order, 0)
	for id, ant := range request.Ants {
		if ok, order := tryUnload(ant); ok {
			orders[id] = order
			continue
		}

		if ok, order := tryConsume(ant); ok {
			orders[id] = order
			continue
		}

		if ok, order := tryMove(ant); ok {
			orders[id] = order
			continue
		}

		orders[id] = &anthive.Order{
			Action:    anthive.ActionStay,
			Direction: anthive.DirectionDown,
		}
	}

	return orders
}
