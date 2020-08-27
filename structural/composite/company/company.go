package company

import (
	"fmt"
	"strings"
)

type ICompany interface {
	Display(depth int)
	LineOfDuty()
}

type ConcreteCompany struct {
	Name string
	List map[string]ICompany
}

func NewConcreteCompany(name string) *ConcreteCompany {
	list := make(map[string]ICompany)
	return &ConcreteCompany{
		Name: name,
		List: list,
	}
}

func (cc *ConcreteCompany) Add(name string, ic ICompany) {
	cc.List[name] = ic
}

func (cc *ConcreteCompany) Remove(name string) {
	delete(cc.List, name)
}

func (cc *ConcreteCompany) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth), " ", cc.Name)
	for _, ccc := range cc.List {
		ccc.Display(depth + 2)
	}
}

func (cc *ConcreteCompany) LineOfDuty() {
	for _, ccc := range cc.List {
		ccc.LineOfDuty()
	}
}
