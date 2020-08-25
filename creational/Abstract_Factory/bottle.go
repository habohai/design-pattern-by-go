package abstract_factory

import "fmt"

type Bottle interface {
	GetBottle()
}

type CocaColaBottle struct{}

func (c *CocaColaBottle) GetBottle() {
	fmt.Println("可口可乐瓶")
}

type PesiColaBottle struct{}

func (c *PesiColaBottle) GetBottle() {
	fmt.Println("百事可乐瓶")
}
