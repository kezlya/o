package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hive struct {
	Username string
	Ants     map[int]*Ant
	Map      *Map
}

type Map struct {
	Width  uint8
	Height uint8
	Cells  [][]*Cell
}

type Cell struct {
	Food uint8
	Ant  string
	Hive string
}

type Ant struct {
	Wasted  int
	Age     uint8
	Health  uint8
	Payload uint8
	X       uint8
	Y       uint8
	Event   string
}

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
