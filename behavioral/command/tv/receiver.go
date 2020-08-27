package tv

import "fmt"

type TV struct{}

func (tv TV) Open() {
	fmt.Println("开机")
}

func (tv TV) Close() {
	fmt.Println("关机")
}

func (tv TV) ChangeChannel() {
	fmt.Println("换台")
}
