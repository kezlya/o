package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"net/http"
)

func StartServer() {
	h := requestHandler

	if err := fasthttp.ListenAndServe(":7070", h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}

}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("content-type; application/json")

	//fmt.Println("Connection at ", ctx.ConnTime())
	//fmt.Println("Request has been started at ", ctx.Time())
	//fmt.Println("Serial request number for the current connection is", ctx.ConnRequestNum())

	body := ctx.Request.Body()

	var hive Hive
	err := json.Unmarshal(body, &hive)
	if err != nil {
		fmt.Println("Fail to convrt json to Object", err)
		return
	}

	hive.Map.getObjects(hive.Id)
	actions := whatToDo(&hive)

	output, err := json.Marshal(actions)
	if err != nil {
		return
	}
	ctx.Write(output)
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

	hive.Map.getObjects(hive.Id)
	actions := whatToDo(&hive)

	output, err := json.Marshal(actions)
	if err != nil {
		http.Error(w, err.Error(), 504)
		return
	}

	fmt.Print("Tick:", hive.Tick, " ")
	fmt.Println(string(output))

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
