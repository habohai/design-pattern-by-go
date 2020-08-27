package composite

import "fmt"

type Leaf struct {
	value int
}

func NewLeaf(value int) *Leaf {
	return &Leaf{value}
}

func (l *Leaf) Traverse() {
	fmt.Println(l.value)
}
