package manager

import "fmt"

type Request struct {
	RequestType    string
	RequestContent string
	Number         int
}

type Manager interface {
	SetNext(next Manager)
	RequestHandler(request Request)
}

type CommonManager struct {
	Manager
	Name string
}

func (cm *CommonManager) SetNext(next Manager) {
	cm.Manager = next
}

func (cm *CommonManager) RequestHandler(request Request) {
	if request.RequestType == "请假" && request.Number <= 2 {
		fmt.Printf("%s: %s 数量 %d 已批准\n", cm.Name, request.RequestContent, request.Number)
	} else {
		if cm.Manager != nil {
			cm.Manager.RequestHandler(request)
		}
	}
}

// MajorManager 主管
type MajorManager struct {
	Manager
	Name string
}

func (mm *MajorManager) SetNext(next Manager) {
	mm.Manager = next
}

func (mm *MajorManager) RequestHandler(request Request) {
	if request.RequestType == "请假" && request.Number <= 5 {
		fmt.Printf("%s: %s 数量 %d 已批准\n", mm.Name, request.RequestContent, request.Number)
	} else {
		if mm.Manager != nil {
			mm.Manager.RequestHandler(request)
		}
	}
}

// GeneralManager 总经理
type GeneralManager struct {
	Manager
	Name string
}

func (gm *GeneralManager) SetNext(next Manager) {
	gm.Manager = next
}

func (gm *GeneralManager) RequestHandler(request Request) {
	if request.RequestType == "请假" {
		fmt.Printf("%s: %s 数量 %d 已批准\n", gm.Name, request.RequestContent, request.Number)
	} else if request.RequestType == "加薪" && request.Number <= 500 {
		fmt.Printf("%s: %s 数量 %d 已批准\n", gm.Name, request.RequestContent, request.Number)
	} else if request.RequestType == "加薪" && request.Number > 500 {
		fmt.Printf("%s: %s 数量 %d 再说吧\n", gm.Name, request.RequestContent, request.Number)
	}
}
