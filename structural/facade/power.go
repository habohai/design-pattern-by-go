package facade

import "fmt"

type Power struct{}

func (p Power) PowerOn() {
	fmt.Println("电源通电")
}

func (p Power) PowerOff() {
	fmt.Println("电源下电")
}

type Mainboard struct{}

func (m Mainboard) PowerOn() {
	fmt.Println("主板通电")
}

func (m Mainboard) PowerOff() {
	fmt.Println("主板下电")
}

type CPU struct{}

func (cpu CPU) PowerOn() {
	fmt.Println("CPU通电")
}

func (cpu CPU) PowerOff() {
	fmt.Println("CPU下电")
}
