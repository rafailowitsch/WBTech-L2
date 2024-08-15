package main

type Handler interface {
	SetNext(handler Handler)
	Handle(request int)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) HandleNext(request int) {
	if h.next != nil {
		h.next.Handle(request)
	}
}
