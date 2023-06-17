package main

var id string
var canvas Canvas

func main() {
	StartServer()
}

func whatToDo(request *Request) []Order {
	orders := make([]Order, 0)
	for _, ant := range request.Ants {
		if ok, order := ant.tryUnload(); ok {
			order.AntId = ant.Id
			orders = append(orders, *order)
			continue
		}

		if ok, order := ant.tryConsume(); ok {
			order.AntId = ant.Id
			orders = append(orders, *order)
			continue
		}

		if ok, order := tryMove(ant); ok {
			order.AntId = ant.Id
			orders = append(orders, *order)
			continue
		}
		order := Order{
			AntId:     ant.Id,
			Action:    ActionStay,
			Direction: DirectionDown,
		}
		orders = append(orders, order)
	}

	return orders
}
