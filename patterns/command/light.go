package main

import "fmt"

type Light struct {
	Active bool
}

func (l *Light) On() {
	l.Active = true
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	l.Active = false
	fmt.Println("Light is OFF")
}
