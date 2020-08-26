package builder

type Director struct {
	builder Builder
}

func (p Director) Create(name string, arms string) *Character {
	return p.builder.SetName(name).SetArms(arms).Build()
}
