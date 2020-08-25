package simple_factory

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

type Factory struct{}

func (f *Factory) NewCola(s string) Cola {
	switch s {
	case "c":
		return &CocaCola{}
	case "p":
		return &PesiCola{}
	}
	return nil
}
