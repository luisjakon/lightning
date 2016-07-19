package lightning

type Handler interface {
	Handle(*Context)
}

type HandlerFunc func(*Context)

func (hf HandlerFunc) Handle(ctx *Context) {
	hf(ctx)
}
