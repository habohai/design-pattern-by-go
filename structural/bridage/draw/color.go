package draw

import "fmt"

type Color interface {
	Use()
}

type Red struct{}

func (r Red) Use() {
	fmt.Println("Use Red color")
}

type Green struct{}

func (g Green) Use() {
	fmt.Println("Use Green color")
}

type Yellow struct{}

func (y Yellow) Use() {
	fmt.Println("Use Yellow color")
}
