package work

import "testing"

func TestWork(t *testing.T) {
	tWork := Work{}
	tState := ForenoonState{}
	tWork.SetState(tState)
	tWork.SetFinishState(true)
	tWork.SetHour(13)
	tWork.WriteProgram()
}
