package lightning

// Middleware Interface.
type Middleware interface {
	Handle(next Handler) Handler // handle request.
}

type BaseMiddleware struct {
}

func (bm *BaseMiddleware) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Invoke the next middleware.
		next.Handle(ctx)
	})
}
