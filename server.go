package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hive struct {
	Id   string
	Ants map[int]*Ant
	Map  *Map
}

type Map struct {
	Width  int
	Height int
	Cells  [][]*Cell
}

type point struct {
	x int
	y int
}

type Cell struct {
	Food int    `json:"food,omitempty"`
	Hive string `json:"hive,omitempty"`
	Ant  string `json:"ant,omitempty"`
}

type Ant struct {
	Wasted  int
	Age     int
	Health  int
	Payload int
	X       int
	Y       int
	Event   string
}

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

func StartServer() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	if len(data) == 0 {
		http.Error(w, err.Error(), 502)
		return
	}

	var hive Hive
	//fmt.Println(string(data))
	err = json.Unmarshal(data, &hive)
	if err != nil {
		fmt.Println("Fail to convrt json to object", err)
		http.Error(w, err.Error(), 503)
		return
	}

	actions := whatToDo(&hive)

	output, err := json.Marshal(actions)
	if err != nil {
		http.Error(w, err.Error(), 504)
		return
	}
	fmt.Println(string(output))

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
