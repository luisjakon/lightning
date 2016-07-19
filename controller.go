package lightning

type ControllerInterface interface {
	Handle(next Handler) Handler

	DELETE(ctx *Context)
	GET(ctx *Context)
	HEAD(ctx *Context)
	OPTIONS(ctx *Context)
	PATCH(ctx *Context)
	POST(ctx *Context)
	PUT(ctx *Context)
}

type Controller struct {
}

func (c Controller) Handle(next Handler) Handler {
	return HandlerFunc(func(ctx *Context) {
		// Invoke the request handler.
		next.Handle(ctx)
	})
}

func (c Controller) DELETE(ctx *Context) {
	ctx.ResponseForbidden()
}

func (c Controller) GET(ctx *Context) {
	ctx.ResponseForbidden()
}

func (c Controller) HEAD(ctx *Context) {
	ctx.ResponseForbidden()
}

func (c Controller) OPTIONS(ctx *Context) {
	ctx.ResponseForbidden()
}

func (c Controller) PATCH(ctx *Context) {
	ctx.ResponseForbidden()
}

func (c Controller) POST(ctx *Context) {
	ctx.ResponseForbidden()
}

func (c Controller) PUT(ctx *Context) {
	ctx.ResponseForbidden()
}
