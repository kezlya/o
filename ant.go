package main

import "github.com/kezlya/anthive"

func tryUnload(a *anthive.Ant) (bool, *anthive.Order) {
	if a.Payload > 0 && a.Y > 0 &&
		canvas.Cells[a.Y-1][a.X].Hive == id &&
		canvas.Cells[a.Y-1][a.X].Ant == "" {
		return true, &anthive.Order{
			Action:    anthive.ActionUnload,
			Direction: anthive.DirectionUp}
	}

	if a.Payload > 0 && a.X < canvas.Width-1 &&
		canvas.Cells[a.Y][a.X+1].Hive == id &&
		canvas.Cells[a.Y][a.X+1].Ant == "" {
		return true, &anthive.Order{
			Action:    anthive.ActionUnload,
			Direction: anthive.DirectionRight}
	}

	if a.Payload > 0 && a.Y < canvas.Height-1 &&
		canvas.Cells[a.Y+1][a.X].Hive == id &&
		canvas.Cells[a.Y+1][a.X].Ant == "" {
		return true, &anthive.Order{
			Action:    anthive.ActionUnload,
			Direction: anthive.DirectionDown}
	}

	if a.Payload > 0 && a.X > 0 &&
		canvas.Cells[a.Y][a.X-1].Hive == id &&
		canvas.Cells[a.Y][a.X-1].Ant == "" {
		return true, &anthive.Order{
			Action:    anthive.ActionUnload,
			Direction: anthive.DirectionLeft}
	}

	return false, nil
}

func tryConsume(a *anthive.Ant) (bool, *anthive.Order) {
	order := &anthive.Order{}

	if a.Health < 9 {
		order.Action = anthive.ActionEat
	} else if a.Payload < 9 {
		order.Action = anthive.ActionLoad
	} else {
		return false, nil
	}

	if isFood(int(a.Y)-1, int(a.X), order) {
		order.Direction = anthive.DirectionUp
		return true, order
	}

	if isFood(int(a.Y), int(a.X)+1, order) {
		order.Direction = anthive.DirectionRight
		return true, order
	}

	if isFood(int(a.Y)+1, int(a.X), order) {
		order.Direction = anthive.DirectionDown
		return true, order
	}

	if isFood(int(a.Y), int(a.X)-1, order) {
		order.Direction = anthive.DirectionLeft
		return true, order
	}

	return false, nil
}

func tryMove(a *anthive.Ant) (bool, *anthive.Order) {
	objects := getObjects()
	shortest := 9999999
	var firstTarget *Object
	var secondTarget *Object
	for _, object := range objects {
		if a.Payload == 9 && !object.hive { // move home
			continue
		}

		if a.Payload < 5 && object.hive { // search for food
			continue
		}

		s := object.distance(int(a.Y), int(a.X))
		if s == 0 {
			continue
		}

		if s < shortest {
			shortest = s
			if !object.used {
				secondTarget = firstTarget
				firstTarget = object
			} else {
				secondTarget = object
			}
		}
	}

	if firstTarget != nil {
		return chooseDirection(a, firstTarget.y, firstTarget.x)
	}

	if secondTarget != nil {
		return chooseDirection(a, secondTarget.y, secondTarget.x)
	}

	return false, nil
}

//TODO: check for future occupied cells by my ants
func chooseDirection(a *anthive.Ant, dy, dx int) (bool, *anthive.Order) {
	if int(a.X) < dx && isEmpty(int(a.Y), int(a.X)+1) {
		return true, &anthive.Order{
			Action:    anthive.ActionMove,
			Direction: anthive.DirectionRight}
	}

	if int(a.Y) < dy && isEmpty(int(a.Y)+1, int(a.X)) {
		return true, &anthive.Order{
			Action:    anthive.ActionMove,
			Direction: anthive.DirectionDown}
	}

	if int(a.X) > dx && isEmpty(int(a.Y), int(a.X)-1) {
		return true, &anthive.Order{
			Action:    anthive.ActionMove,
			Direction: anthive.DirectionLeft}
	}

	if int(a.Y) > dy && isEmpty(int(a.Y)-1, int(a.X)) {
		return true, &anthive.Order{
			Action:    anthive.ActionMove,
			Direction: anthive.DirectionUp}
	}

	return false, nil
}
