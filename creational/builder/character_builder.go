package builder

type CharacterBuilder struct {
	character *Character
}

func (c *CharacterBuilder) SetName(name string) Builder {
	if c.character == nil {
		c.character = &Character{}
	}
	c.character.SetName(name)
	return c
}

func (c *CharacterBuilder) SetArms(arms string) Builder {
	if c.character == nil {
		c.character = &Character{}
	}
	c.character.SetArms(arms)
	return c
}

func (c *CharacterBuilder) Build() *Character {
	return c.character
}
