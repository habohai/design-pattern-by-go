package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	cf := &CocaFactory{}
	cf.NewCola().GetCola()
	cf.NewBottle().GetBottle()
	cf.NewBox().GetBox()

	pf := &PesiFactory{}
	pf.NewCola().GetCola()
	pf.NewBottle().GetBottle()
	pf.NewBox().GetBox()
}
