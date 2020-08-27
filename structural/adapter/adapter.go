package adapter

type Adapter struct {
	Adaptee
}

func (a Adapter) Request() {
	a.SpecificRequest()
}

func NewAdapter(adaptee Adaptee) Adapter {
	return Adapter{
		Adaptee: adaptee,
	}
}
