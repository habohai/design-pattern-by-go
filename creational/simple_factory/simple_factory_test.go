package simple_factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	f := &Factory{}
	f.NewCola("p").GetCola()
	f.NewCola("c").GetCola()
}
