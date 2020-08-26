package decorator

type Sausage struct {
	noddles Noddles
	name    string
	price   float32
}

func (p Sausage) SetNoddles(noddles Noddles) {
	p.noddles = noddles
}

func (p Sausage) Description() string {
	return p.noddles.Description() + "+" + p.name
}

func (p Sausage) Price() float32 {
	return p.noddles.Price() + p.price
}
