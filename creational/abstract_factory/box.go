package abstract_factory

import "fmt"

type Box interface {
	GetBox()
}

type CocaColaBox struct{}

func (c *CocaColaBox) GetBox() {
	fmt.Println("可口可乐箱")
}

type PesiColaBox struct{}

func (c *PesiColaBox) GetBox() {
	fmt.Println("百事可乐箱")
}
