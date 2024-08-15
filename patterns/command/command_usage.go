package main

func CommandUsage() {
	// Получатель
	light := &Light{}

	// Команды
	lightOn := &LightOnCommand{light}
	lightOff := &LightOffCommand{light}

	// Инициатор
	remote := &RemoteControl{}

	// Клиент
	remote.SetCommand(lightOn)
	remote.SetCommand(lightOff)

	// Отмена последней команды
	remote.Undo()
	remote.Undo()
}

func main() {
	CommandUsage()
}
