package boss

type Subject interface {
	SetState(state string)
	GetState() string
	Attach(observer Observer)
	Notify()
}

type Boss struct {
	observer []Observer
	state    string
}

func (b *Boss) SetState(state string) {
	b.state = state
}

func (b Boss) GetState() string {
	return b.state
}

func (b *Boss) Attach(observer Observer) {
	b.observer = append(b.observer, observer)
}

func (b Boss) Notify() {
	for _, o := range b.observer {
		o.Update()
	}
}

type Secretary struct {
	observer []Observer
	state    string
}

func (s *Secretary) SetState(state string) {
	s.state = state
}

func (s Secretary) GetState() string {
	return s.state
}

func (s *Secretary) Attach(observer Observer) {
	s.observer = append(s.observer, observer)
}

func (s Secretary) Notify() {
	for _, o := range s.observer {
		o.Update()
	}
}

func NewSubject(t string) Subject {
	switch t {
	case "boss":
		return &Boss{}
	case "secretary":
		return &Secretary{}
	default:
		return nil
	}
}
