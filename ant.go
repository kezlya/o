package main

func tryUnload(a Ant) (bool, *Order) {
	if a.Payload > 0 && a.Point.Y > 0 &&
		canvas.Cells[a.Point.Y-1][a.Point.X].Hive == id &&
		canvas.Cells[a.Point.Y-1][a.Point.X].Ant == "" {
		return true, &Order{
			Action:    ActionPut,
			Direction: DirectionUp}
	}

	if a.Payload > 0 && a.Point.X < canvas.Width-1 &&
		canvas.Cells[a.Point.Y][a.Point.X+1].Hive == id &&
		canvas.Cells[a.Point.Y][a.Point.X+1].Ant == "" {
		return true, &Order{
			Action:    ActionPut,
			Direction: DirectionRight}
	}

	if a.Payload > 0 && a.Point.Y < canvas.Height-1 &&
		canvas.Cells[a.Point.Y+1][a.Point.X].Hive == id &&
		canvas.Cells[a.Point.Y+1][a.Point.X].Ant == "" {
		return true, &Order{
			Action:    ActionPut,
			Direction: DirectionDown}
	}

	if a.Payload > 0 && a.Point.X > 0 &&
		canvas.Cells[a.Point.Y][a.Point.X-1].Hive == id &&
		canvas.Cells[a.Point.Y][a.Point.X-1].Ant == "" {
		return true, &Order{
			Action:    ActionPut,
			Direction: DirectionLeft}
	}

	return false, nil
}

func tryConsume(a Ant) (bool, *Order) {
	order := &Order{}

	if a.Health < 9 {
		order.Action = ActionEat
	} else if a.Payload < 9 {
		order.Action = ActionTake
	} else {
		return false, nil
	}

	if isFood(a.Point.Y-1, a.Point.X, order) {
		order.Direction = DirectionUp
		return true, order
	}

	if isFood(a.Point.Y, a.Point.X+1, order) {
		order.Direction = DirectionRight
		return true, order
	}

	if isFood(a.Point.Y+1, a.Point.X, order) {
		order.Direction = DirectionDown
		return true, order
	}

	if isFood(a.Point.Y, a.Point.X-1, order) {
		order.Direction = DirectionLeft
		return true, order
	}

	return false, nil
}

func tryMove(a Ant) (bool, *Order) {
	objects := getObjects()
	var shortest uint = 9999999
	var firstTarget *Object
	var secondTarget *Object
	for _, object := range objects {
		if a.Payload == 9 && !object.hive { // move home
			continue
		}

		if a.Payload < 5 && object.hive { // search for food
			continue
		}

		s := object.distance(a.Point.Y, a.Point.X)
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
func chooseDirection(a Ant, dy, dx uint) (bool, *Order) {
	if a.Point.X < dx && isEmpty(a.Point.Y, a.Point.X+1) {
		return true, &Order{
			Action:    ActionMove,
			Direction: DirectionRight}
	}

	if a.Point.Y < dy && isEmpty(a.Point.Y+1, a.Point.X) {
		return true, &Order{
			Action:    ActionMove,
			Direction: DirectionDown}
	}

	if a.Point.X > dx && isEmpty(a.Point.Y, a.Point.X-1) {
		return true, &Order{
			Action:    ActionMove,
			Direction: DirectionLeft}
	}

	if a.Point.Y > dy && isEmpty(a.Point.Y-1, a.Point.X) {
		return true, &Order{
			Action:    ActionMove,
			Direction: DirectionUp}
	}

	return false, nil
}
