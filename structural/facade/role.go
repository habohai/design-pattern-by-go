package facade

import "fmt"

type Computer interface {
	Boot()
	Shutdown()
}

func NewComputer(s string) Computer {
	switch s {
	case "d":
		return &DesktopComputer{}
	case "l":
		return &Laptop{}
	default:
		return nil
	}
}

type DesktopComputer struct {
	Power
	Mainboard
	CPU
}

func (dc DesktopComputer) Boot() {
	fmt.Println("台式机开始启动")
	dc.Power.PowerOn()
	dc.Mainboard.PowerOn()
	dc.CPU.PowerOn()
	fmt.Println("台式机启动完毕")
}

func (dc DesktopComputer) Shutdown() {
	fmt.Println("台式机开始关机")
	dc.CPU.PowerOff()
	dc.Mainboard.PowerOff()
	dc.Power.PowerOff()
	fmt.Println("台式机关机完毕")
}

type Laptop struct {
	Power
	Mainboard
	CPU
}

func (l Laptop) Boot() {
	fmt.Println("笔记本电脑开始启动")
	l.Power.PowerOn()
	l.Mainboard.PowerOn()
	l.CPU.PowerOn()
	fmt.Println("笔记本电脑启动完毕")
}

func (l Laptop) Shutdown() {
	fmt.Println("笔记本电脑开始关机")
	l.CPU.PowerOff()
	l.Mainboard.PowerOff()
	l.Power.PowerOff()
	fmt.Println("笔记本电脑关机完毕")
}
