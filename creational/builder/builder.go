package builder

type Builder interface {
	SetName(name string) Builder
	SetArms(arms string) Builder
	Build() *Character
}

