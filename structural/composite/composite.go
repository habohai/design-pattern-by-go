package composite

type Composite struct {
	children []Component
}

func NewComposite() *Composite {
	return &Composite{children: make([]Component, 0)}
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Traverse() {
	for idx, _ := range c.children {
		c.children[idx].Traverse()
	}
}
