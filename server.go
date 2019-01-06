package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	err = json.Unmarshal(data, &hive)
	if err != nil {
		fmt.Println("Fail to convrt json to Object", err)
		http.Error(w, err.Error(), 503)
		return
	}

	hive.allFood = hive.Map.getObjects(hive.Id)
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
