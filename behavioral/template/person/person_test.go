package person

import (
	"fmt"
	"testing"
)

func TestPerson(t *testing.T) {
	person := &Person{}
	person.Specific = &Boy{}
	person.SetName("zhangshan")
	person.Out()
	fmt.Println("----------------------")
	person.Specific = &Girl{}
	person.SetName("lisi")
	person.Out()
}
