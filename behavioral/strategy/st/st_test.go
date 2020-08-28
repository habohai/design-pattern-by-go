package st

import "testing"

func TestSt(t *testing.T) {
	tStrategy := NewStrategy("A")
	tContext := NewContext(tStrategy)
	tContext.Algorithm()

	tStrategy = NewStrategy("A")
	tContext = NewContext(tStrategy)
	tContext.Algorithm()

	tStrategy = NewStrategy("C")
	tContext = NewContext(tStrategy)
	tContext.Algorithm()
}
