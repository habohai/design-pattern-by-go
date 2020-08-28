package st

type Context struct {
	strategy Strategy
}

func (c Context) Algorithm() {
	c.strategy.AlgorithmInterface()
}

func NewContext(s Strategy) Context {
	return Context{
		strategy: s,
	}
}
