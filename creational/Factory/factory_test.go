package factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	// 根据不同的工厂类生产不同的产品
	c := &CocaFactory{}
	c.NewCola().GetCola()
	p := &PesiFactory{}
	p.NewCola().GetCola()
}
