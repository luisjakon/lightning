package middleware

import (
	"bytes"
	"github.com/headwindfly/jwt"
	"github.com/luisjakon/lightning"
)

const (
	urlKey  = "_jwt"
	formKey = "_jwt"
)

var (
	bearer = []byte("BEARER ")
)

type JWTMiddleware struct {
	jwt     *jwt.JWT
	urlKey  string
	formKey string
}

func NewJWTMiddleware(jwt *jwt.JWT) JWTMiddleware {
	return JWTMiddleware{
		jwt:     jwt,
		urlKey:  urlKey,
		formKey: formKey,
	}
}

func (jm JWTMiddleware) Handle(next lightning.Handler) lightning.Handler {
	return lightning.HandlerFunc(func(ctx *lightning.Context) {
		// Try to get JWT raw token from URL query string.
		rawToken := ctx.FormValue(jm.urlKey)
		if len(rawToken) == 0 {
			// Try to get JWT raw token from POST FORM.
			rawToken = ctx.FormValue(jm.formKey)
			if len(rawToken) == 0 {
				// Try to get JWT raw token from Header.
				if ah := ctx.Request.Header.Peek("Authorization"); len(ah) > 0 {
					// Should be a bearer token
					if len(ah) > 6 && bytes.Equal(ah[:7], bearer) {
						rawToken = ah[7:]
					}
				}
			}
		}

		// Check raw token is valid.
		if len(rawToken) == 0 {
			ctx.ResponseUnauthorized()
			return
		}

		// Get JWT by raw token.
		var err error
		ctx.Token, err = jm.jwt.NewTokenByRaw(string(rawToken))
		if err != nil {
			ctx.ResponseUnauthorized()
			return
		}

		// Validate Token.
		if err := ctx.Token.Validate(); err != nil {
			ctx.ResponseUnauthorized()
			return
		}

		// Validate successfully.
		next.Handle(ctx)
	})
}
