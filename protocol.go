package main

type Request struct {
	Id     string `json:"id"`
	Tick   uint   `json:"tick"`
	Ants   []Ant  `json:"ants"`
	Canvas Canvas `json:"canvas"`
}

type Response struct {
	Orders []Order `json:"orders"`
}

type Ant struct {
	Id      uint   `json:"id"`
	Event   string `json:"event"`
	Errors  uint   `json:"errors"`
	Age     uint   `json:"age"`
	Health  uint   `json:"health"`
	Payload uint   `json:"payload"`
	Point   Point  `json:"point"`
}

type Order struct {
	AntId     uint   `json:"antId"`
	Action    string `json:"act"`
	Direction string `json:"dir"`
}

const ActionStay = "stay"
const ActionMove = "move"
const ActionTake = "take"
const ActionPut = "put"
const ActionEat = "eat"

const DirectionUp = "up"
const DirectionRight = "right"
const DirectionDown = "down"
const DirectionLeft = "left"

const EventGood = "good"
const EventBirth = "birth"
const EventNoAction = "noaction"
const EventSlow = "slow"
const EventBadMove = "badmove"
const EventBadLoad = "badload"
const EventBadUnload = "badunload"
const EventBadEat = "badeat"
const EventCollision = "collision"
const EventError = "error"
const EventDeath = "death"

type Map struct {
	Canvas Canvas `json:"canvas"`
	Food   uint   `json:"food"`
	Hunger uint   `json:"hunger"`
	Theme  uint   `json:"theme"`
}

type Canvas struct {
	Width  uint     `json:"width"`
	Height uint     `json:"height"`
	Cells  [][]Cell `json:"cells"`
}

type Cell struct {
	Food    uint   `json:"food,omitempty"`
	Terrain string `json:"terrain,omitempty"`
	Hive    string `json:"hive,omitempty"`
	Ant     string `json:"ant,omitempty"`
}

type Point struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}
