package work

type Work struct {
	hour    int
	current State
	finish  bool
}

func (w *Work) SetState(s State) {
	w.current = s
}

func (w *Work) SetHour(hour int) {
	w.hour = hour
}

func (w *Work) SetFinishState(finish bool) {
	w.finish = finish
}

func (w *Work) WriteProgram() {
	w.current.WriteProgram(w)
}
