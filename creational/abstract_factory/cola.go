package abstract_factory

import "fmt"

type Cola interface {
	GetCola()
}

type CocaCola struct {
}

func (d *CocaCola) GetCola() {
	fmt.Println("可口可乐")
}

type PesiCola struct {
}

func (q *PesiCola) GetCola() {
	fmt.Println("百事可乐")
}
