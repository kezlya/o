package main

import "github.com/kezlya/anthive"

var id string
var canvas anthive.Canvas

func main() {
	StartServer()
}

func whatToDo(request *anthive.Request) []anthive.Order {
	orders := make([]anthive.Order, 0)
	for _, ant := range request.Ants {
		if ok, order := tryUnload(ant); ok {
			order.AntId = ant.Id
			orders = append(orders, *order)
			continue
		}

		if ok, order := tryConsume(ant); ok {
			order.AntId = ant.Id
			orders = append(orders, *order)
			continue
		}

		if ok, order := tryMove(ant); ok {
			order.AntId = ant.Id
			orders = append(orders, *order)
			continue
		}
		order := anthive.Order{
			AntId:     ant.Id,
			Action:    anthive.ActionStay,
			Direction: anthive.DirectionDown,
		}
		orders = append(orders, order)
	}

	return orders
}
