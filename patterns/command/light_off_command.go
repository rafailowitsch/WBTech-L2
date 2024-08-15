package main

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}
