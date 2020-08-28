package person

import "fmt"

type Boy struct {
	Person
}

func (b *Boy) BeforeOut() {
	fmt.Println("get up...")
}

type Girl struct {
	Person
}

func (g *Girl) BeforeOut() {
	fmt.Println("get up and dress up...")
}
