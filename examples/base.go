// Copyright 2016 Bolt. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/luisjakon/lightning"
	"log"
	"os"
	"path"
)

var (
	helloLightning = []byte("Hello Lightning!\n")
	resourcesPath  = path.Join(os.Getenv("GOPATH"), "src", "github.com", "luisjakon", "lightning", "examples")
)

type User struct {
	Name string `json:"name" xml:"name"`
	Team string `json:"team" xml:"team"`
}

func hello(ctx *lightning.Context) {
	ctx.Write(helloLightning)
}

func html(ctx *lightning.Context) {
	ctx.HTML("Hello Lightning!\n")
}

func json(ctx *lightning.Context) {
	ctx.JSON(User{
		Name: "Bolt",
		Team: "Lightning",
	})
}

func jsonp(ctx *lightning.Context) {
	callback := ctx.FormValue("callback")
	ctx.JSONP(User{
		Name: "Bolt",
		Team: "Lightning",
	}, callback)
}

func xml(ctx *lightning.Context) {
	ctx.XML(User{
		Name: "Bolt",
		Team: "Lightning",
	}, "")
}

func params(ctx *lightning.Context) {
	name := ctx.RouterParams.ByName("name")
	ctx.Textf("Hello %s.", name)
}

func multiParams(ctx *lightning.Context) {
	param1 := ctx.Params("param1")
	param2 := ctx.Params("param2")
	ctx.Textf("Your params is %s and %s", param1, param2)
}

func main() {
	// Create a router instance.
	router := lightning.NewRouter()

	// Register route handler.
	router.GET("/", lightning.HandlerFunc(hello))
	router.GET("/html", lightning.HandlerFunc(html))
	router.GET("/json", lightning.HandlerFunc(json))
	router.GET("/jsonp", lightning.HandlerFunc(jsonp))
	router.GET("/xml", lightning.HandlerFunc(xml))
	// Navigate to http://127.0.0.1:8080/params/yourname.
	router.GET("/params/:name", lightning.HandlerFunc(params))
	// Navigate to http://127.0.0.1:8080/multi-params/param1/param2.
	router.GET("/multi-params/:param1/:param2", lightning.HandlerFunc(multiParams))

	// Static resource files.
	// Navigate to http://127.0.0.1:8080/examples/base.go
	router.ServeFiles("/examples/*filepath", resourcesPath)

	// Start server.
	log.Fatal(lightning.ListenAndServe(":8080", router.Handler))
}
