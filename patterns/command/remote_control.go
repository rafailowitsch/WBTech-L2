package main

import "fmt"

type RemoteControl struct {
	history []Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.history = append(r.history, command)
	command.Execute()
}

func (r *RemoteControl) Undo() {
	if len(r.history) > 0 {
		lastCommand := r.history[len(r.history)-1]
		lastCommand.Undo()
		r.history = r.history[:len(r.history)-1]
	} else {
		fmt.Println("No commands to undo")
	}
}
