// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto"
	"github.com/headwindfly/jwt"
	"github.com/luisjakon/lightning"
	"github.com/luisjakon/lightning/middlewares"
	"log"
	"math/rand"
	"strconv"
)

var (
	ttl = int64(20)
	j   = jwt.NewJWT("Lightning", ttl)
)

func init() {
	// Add HMACAlgorithm.
	hs256, err := jwt.NewHMACAlgorithm(crypto.SHA256, []byte("secrey key"))
	if err != nil {
		panic(err)
	}
	j.AddAlgorithm("HS256", hs256)
}

func jwtGet(ctx *lightning.Context) {
	token, err := j.NewToken("HS256", "Lightning", "audience"+strconv.Itoa(rand.Intn(100)))
	if err != nil {
		ctx.Textf("Failed to generate token: %s", err.Error())
		return
	}
	token.Parse()
	ctx.Textf("JWT token(effective within %d seconds):\n%s", ttl, token.Raw.Token())
}

func jwtVerify(ctx *lightning.Context) {
	ctx.Text("Congratulation! Your token is valid.")
}

func main() {
	// Create a router instance.
	router := lightning.NewRouter()

	// Note that.
	// Before registering middleware, we should register jwtGet handler first.
	// In order to cross over the JWT Middleware,
	// otherwise the jwtGet handler will be blocked by the JWT Middleware.
	router.GET("/", lightning.HandlerFunc(jwtGet))

	// Add JWT Middleware
	router.AddMiddleware(middleware.NewJWTMiddleware(j))

	// Register route handler.
	router.GET("/verify", lightning.HandlerFunc(jwtVerify))
	router.POST("/verify", lightning.HandlerFunc(jwtVerify))

	// Start server.
	log.Fatal(lightning.ListenAndServe(":8080", router.Handler))
}
