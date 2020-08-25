package factory

import "fmt"

type Cola interface {
	GetCola()
}

type Factory interface{
	NewCola() Cola
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

type PesiFactory struct{}

func (p *PesiFactory) NewCola() Cola {
	return &PesiCola{}
}

type CocaFactory struct{}

func (c *CocaFactory) NewCola() Cola{
	return &CocaCola{}
}