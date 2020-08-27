package tv

type Invoke struct {
	Command
}

func (i Invoke) ExecuteCommand() {
	i.Command.Execute()
}
