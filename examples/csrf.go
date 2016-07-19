// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/headwindfly/lightning"
	"github.com/headwindfly/lightning/middlewares"
	"github.com/headwindfly/sessions"
	"log"
	"time"
)

var csrf = middleware.NewCSRFMiddleware()

func CSRFGet(ctx *lightning.Context) {
	ctx.HTML(fmt.Sprintf(`
	CSRF token: <br>
	<b>%s</b><br>
	Send post request to http://127.0.0.1:8080/ to verify the CSRF token.<br>
	You should store the token arg into request's header(<b>%s</b>) or form(<b>%s</b>).<br>
	If it does not work, make sure that your redis-server is started.<br>
	`,
		ctx.UserValue(csrf.Key),
		csrf.HeaderKey,
		csrf.FormKey,
	))
}

func CSRFPost(ctx *lightning.Context) {
	ctx.HTML("Congratulation! Your token is valid.\n")
}

func main() {
	// Create a router instance.
	router := lightning.NewRouter()

	// Set session store.
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

	// Add CSRF middleware.
	router.AddMiddleware(csrf)

	// Register route handler.
	router.GET("/", lightning.HandlerFunc(CSRFGet))
	router.POST("/", lightning.HandlerFunc(CSRFPost))

	// Start server.
	log.Fatal(lightning.ListenAndServe(":8080", router.Handler))
}
