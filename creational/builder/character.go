package builder

// Character 人物类
type Character struct {
	Name string
	Arms string
}

func (p *Character) SetName(name string) {
	p.Name = name
}

func (p *Character) SetArms(arms string) {
	p.Arms = arms
}

func (p Character) GetName() string {
	return p.Name
}

func (p Character) GetArms() string {
	return p.Arms
}
