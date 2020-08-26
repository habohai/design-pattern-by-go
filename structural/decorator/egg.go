package decorator

type Egg struct {
	noddles Noddles
	name    string
	price   float32
}

func (p Egg) SetNoddles(noddles Noddles) {
	p.noddles = noddles
}

func (p Egg) Description() string {
	return p.noddles.Description() + "+" + p.name
}

func (p Egg) Price() float32 {
	return p.noddles.Price() + p.price
}
