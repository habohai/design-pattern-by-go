package work

import "fmt"

type State interface {
	WriteProgram(work *Work)
}

type ForenoonState struct{}

func (fs ForenoonState) WriteProgram(work *Work) {
	if work.hour < 12 {
		fmt.Println("上午")
	} else {
		work.SetState(NoonState{})
		work.WriteProgram()
	}
}

type NoonState struct{}

func (ns NoonState) WriteProgram(work *Work) {
	if work.hour < 13 {
		fmt.Println("中午")
	} else {
		work.SetState(AfternoonState{})
		work.WriteProgram()
	}
}

type AfternoonState struct{}

func (as AfternoonState) WriteProgram(work *Work) {
	if work.hour < 17 {
		fmt.Println("下午")
	} else {
		work.SetState(EveningState{})
		work.WriteProgram()
	}
}

type EveningState struct{}

func (es EveningState) WriteProgram(work *Work) {
	if work.finish {
		work.SetState(RestState{})
		work.WriteProgram()
	} else {
		if work.hour < 21 {
			fmt.Println("晚间")
		} else {
			work.SetState(SleepingState{})
			work.WriteProgram()
		}
	}
}

type SleepingState struct{}

func (ss SleepingState) WriteProgram(work *Work) {
	fmt.Println("睡着了")
}

type RestState struct{}

func (rs RestState) WriteProgram(work *Work) {
	fmt.Println("下班回家")
}
