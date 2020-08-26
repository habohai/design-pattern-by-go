package facade

import (
	"testing"
)

func TestCarFacade_CreateCompleteCar(t *testing.T) {
	facade := NewCarFacade()
	facade.CreateCompleteCar()
}

func TestComputer(t *testing.T) {
	tComputer := NewComputer("d")
	tComputer.Boot()
	t.Log("=============")
	tComputer.Shutdown()
	t.Log("=============")
	tComputer = NewComputer("l")
	tComputer.Boot()
	t.Log("=============")
	tComputer.Shutdown()
}
