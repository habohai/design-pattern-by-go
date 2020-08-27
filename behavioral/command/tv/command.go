package tv

type Command interface {
	Execute()
}

type OpenCommand struct {
	receiver TV
}

func (oc OpenCommand) Execute() {
	oc.receiver.Open()
}

type CloseCommand struct {
	receiver TV
}

func (cc CloseCommand) Execute() {
	cc.receiver.Close()
}

type ChangeChannelCommand struct {
	receiver TV
}

func (ccc ChangeChannelCommand) Execute() {
	ccc.receiver.ChangeChannel()
}

func NewCommand(t string, tv TV) Command {
	switch t {
	case "open":
		return OpenCommand{
			receiver: tv,
		}
	case "close":
		return CloseCommand{
			receiver: tv,
		}
	case "changechannel":
		return ChangeChannelCommand{
			receiver: tv,
		}
	default:
		return nil
	}
}
