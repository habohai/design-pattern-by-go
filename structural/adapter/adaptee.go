package adapter

import "fmt"

type Adaptee interface {
	SpecificRequest()
}

type XiaoMiAdaptee struct{}

func (xma XiaoMiAdaptee) SpecificRequest() {
	fmt.Println("This is a specific request from Xiao Mi")
}

type HuaWeiAdaptee struct{}

func (hwa HuaWeiAdaptee) SpecificRequest() {
	fmt.Println("This is a specific request from Hua Wei")
}

func NewAdaptee(t string) Adaptee {
	switch t {
	case "XiaoMi":
		return XiaoMiAdaptee{}
	case "HuaWei":
		return HuaWeiAdaptee{}
	default:
		return nil
	}
}
