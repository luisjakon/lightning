// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/headwindfly/sessions"
	"github.com/luisjakon/lightning"
	"log"
	"math/rand"
	"time"
)

var router *lightning.Router

func getSession(ctx *lightning.Context) {
	// Get session.
	ctx.GetSession()
	defer ctx.SaveSession()

	if number, ok := ctx.Session.Values["randomNumber"]; ok {
		fmt.Fprint(ctx, fmt.Sprintf("The random number is: %d.\n", number))
		return
	}

	fmt.Fprint(ctx, "No random number.\n")

	fmt.Fprintf(ctx, "If it does not work, make sure that your redis-server is started.")
}

func setSession(ctx *lightning.Context) {
	// Get session.
	ctx.GetSession()
	defer ctx.SaveSession()

	// Set random number.
	randomNumber := rand.Intn(100)
	ctx.Session.Values["randomNumber"] = randomNumber

	fmt.Fprint(ctx, fmt.Sprintf("The random number has been set as: %d.\n", randomNumber))
}

func main() {
	// Create a router instance.
	router = lightning.NewRouter()

	// Create a redis pool.
	redisPool := &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: time.Duration(100) * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
	defer redisPool.Close()

	// Create a redis session store.
	store := sessions.NewRedisStore(redisPool, sessions.Options{
		MaxAge: 3600 * 24 * 7, // 10 seconds.
		// Domain:".lightning.dev",
		// HttpOnly:true,
		// Secure:true,
	})

	// Set session store.
	router.SetSessionStore(store)

	// Register route handler.
	router.GET("/", lightning.HandlerFunc(getSession))
	router.GET("/random", lightning.HandlerFunc(setSession))

	// Start server.
	log.Fatal(lightning.ListenAndServe(":8080", router.Handler))
}
