package main

import "fmt"

type LowLevelSupport struct {
	BaseHandler
}

func (h *LowLevelSupport) Handle(request int) {
	if request <= 1 {
		fmt.Println("Low-Level Support handles request", request)
	} else {
		fmt.Println("Low-Level Support passes request", request)
		h.HandleNext(request)
	}
}
