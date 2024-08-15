package main

import "fmt"

type HighLevelSupport struct {
	BaseHandler
}

func (h *HighLevelSupport) Handle(request int) {
	if request <= 3 {
		fmt.Println("High-Level Support handles request", request)
	} else {
		fmt.Println("Request", request, "cannot be handled by High-Level Support")
	}
}
