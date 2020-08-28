package st

import "fmt"

type Strategy interface {
	AlgorithmInterface()
}

type StrategyA struct{}

func (sa StrategyA) AlgorithmInterface() {
	fmt.Println("算法A")
}

type StrategyB struct{}

func (sb StrategyB) AlgorithmInterface() {
	fmt.Println("算法B")
}

type StrategyC struct{}

func (sc StrategyC) AlgorithmInterface() {
	fmt.Println("算法C")
}

func NewStrategy(t string) Strategy {
	switch t {
	case "A":
		return StrategyA{}
	case "B":
		return StrategyB{}
	case "C":
		return StrategyC{}
	default:
		return nil
	}
}
