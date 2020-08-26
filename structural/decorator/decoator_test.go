package decorator

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPi(t *testing.T) {
	fmt.Println(Pi(5000))
	fmt.Println(Pi(10000))
	fmt.Println(Pi(50000))

}

func TestWrapLogger(t *testing.T) {
	f := WrapLogger(Pi, log.New(os.Stdout, "test ", 1))
	f(100000)
}

func TestNoddle(t *testing.T) {
	ramen := Ramen{name: "ramen", price: 10}
	t.Log(ramen.Description())
	t.Log(ramen.Price())

	egg := Egg{noddles: ramen, name: "egg", price: 2}

	t.Log(egg.Description())
	t.Log(egg.Price())

	sausage := Sausage{noddles: egg, name: "sausage", price: 3}
	t.Log(sausage.Description())
	t.Log(sausage.Price())

	egg2 := Egg{noddles: egg, name: "egg", price: 2}

	t.Log(egg2.Description())
	t.Log(egg2.Price())
}
