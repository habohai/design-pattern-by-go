package abstract_factory

type Factory interface {
	NewCola() Cola
	NewBottle() Bottle
	NewBox() Box
}

type PesiFactory struct{}

func (p *PesiFactory) NewCola() Cola {
	return &PesiCola{}
}

func (p *PesiFactory) NewBottle() Bottle {
	return &PesiColaBottle{}
}

func (p *PesiFactory) NewBox() Box {
	return &PesiColaBox{}
}

type CocaFactory struct{}

func (c *CocaFactory) NewCola() Cola {
	return &CocaCola{}
}

func (c *CocaFactory) NewBottle() Bottle {
	return &CocaColaBottle{}
}

func (c *CocaFactory) NewBox() Box {
	return &CocaColaBox{}
}
