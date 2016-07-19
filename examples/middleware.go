// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/lightning/lightning"
	"log"
)

// First Middleware
type FirstMiddleware struct {
}

func (fm *FirstMiddleware) Handle(next lightning.Handler) lightning.Handler {
	return lightning.HandlerFunc(func(ctx *lightning.Context) {
		fmt.Fprint(ctx, "I am First Middleware!\n")
		// Invoke the next middleware
		next.Handle(ctx)
	})
}

// Second Middleware
type SecondMiddleware struct {
}

func (sm *SecondMiddleware) Handle(next lightning.Handler) lightning.Handler {
	return lightning.HandlerFunc(func(ctx *lightning.Context) {
		fmt.Fprint(ctx, "I am Second Middleware!\n")
		// Invoke the next middleware
		next.Handle(ctx)
	})
}

func middleware(ctx *lightning.Context) {
	fmt.Fprint(ctx, "Hello Lightning!\n")
}

func main() {
	// Create a router instance.
	router := lightning.NewRouter()

	// Add middleware before registering route's handler.
	router.AddMiddleware(&FirstMiddleware{})
	router.AddMiddleware(&SecondMiddleware{})

	// Register route handler.
	router.GET("/", lightning.HandlerFunc(middleware))

	// Start server.
	log.Fatal(lightning.ListenAndServe(":8080", router.Handler))
}
