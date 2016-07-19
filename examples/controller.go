package main

import (
	"github.com/luisjakon/lightning"
	"log"
)

type UserController struct {
	lightning.Controller
}

func (c UserController) Handle(next lightning.Handler) lightning.Handler {
	return lightning.HandlerFunc(func(ctx *lightning.Context) {
		// Do anything what you want.

		ctx.Text("Prepare.\n")

		// Invoke the request handler.
		next.Handle(ctx)

		ctx.Text("Finished.\n")
	})
}

func (c UserController) GET(ctx *lightning.Context) {
	ctx.Text("GET REQUEST.\n")
}

func (c UserController) POST(ctx *lightning.Context) {
	ctx.Text("POST REQUEST.\n")
}

func (c UserController) DELETE(ctx *lightning.Context) {
	ctx.Text("DELETE REQUEST.\n")
}

func (c UserController) PUT(ctx *lightning.Context) {
	ctx.Text("PUT REQUEST.\n")
}

func (c UserController) OPTIONS(ctx *lightning.Context) {
	ctx.Text("OPTIONS REQUEST.\n")
}

func (c UserController) PATCH(ctx *lightning.Context) {
	ctx.Text("PATCH REQUEST.\n")
}

func (c UserController) HEAD(ctx *lightning.Context) {
	ctx.Text("HEAD REQUEST.\n")
}

func main() {
	// Create a router instance.
	router := lightning.NewRouter()

	// Register route handler.
	router.RegisterController("/", UserController{})

	// Start server.
	log.Fatal(lightning.ListenAndServe(":8080", router.Handler))
}
