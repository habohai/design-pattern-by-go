package person

import "fmt"

type IPerson interface {
	SetName(name string)
	BeforeOut()
	Out()
}

type Person struct {
	Specific IPerson
	name     string
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) BeforeOut() {
	if p.Specific == nil {
		return
	}

	p.Specific.BeforeOut()
}

func (p *Person) Out() {
	p.BeforeOut()
	fmt.Println(p.name + " go out...")
}
