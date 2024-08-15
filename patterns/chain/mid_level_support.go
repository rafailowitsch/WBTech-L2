package main

import "fmt"

type MidLevelSupport struct {
	BaseHandler
}

func (h *MidLevelSupport) Handle(request int) {
	if request <= 2 {
		fmt.Println("Mid-Level Support handles request", request)
	} else {
		fmt.Println("Mid-Level Support passes request", request)
		h.HandleNext(request)
	}
}
