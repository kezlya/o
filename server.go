package main

import (
	"encoding/json"
	"fmt"
	"github.com/kezlya/anthive"
	"github.com/valyala/fasthttp"
	"log"
)

func StartServer() {
	if err := fasthttp.ListenAndServe(":7070", requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}

}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("content-type; application/json")
	body := ctx.Request.Body()

	var request anthive.Request
	err := json.Unmarshal(body, &request)
	if err != nil {
		fmt.Println("Fail to convrt json to Object", err)
		return
	}

	id = request.Id
	canvas = request.Canvas
	actions := whatToDo(&request)

	output, err := json.Marshal(actions)
	if err != nil {
		return
	}
	ctx.Write(output)
}
