package boss

import "fmt"

type Observer interface {
	Update()
}

type StockObserver struct {
	name    string
	subject Subject
}

func (so StockObserver) Update() {
	fmt.Printf("%s, %s 关闭炒股软件，开始工作\n", so.subject.GetState(), so.name)
}

type NBAObserver struct {
	name    string
	subject Subject
}

func (no NBAObserver) Update() {
	fmt.Printf("%s, %s 关闭NBA直播，开始工作\n", no.subject.GetState(), no.name)
}

func NewObserver(t, name string, subject Subject) Observer {
	switch t {
	case "stock":
		return &StockObserver{
			name:    name,
			subject: subject,
		}
	case "nba":
		return &NBAObserver{
			name:    name,
			subject: subject,
		}
	default:
		return nil
	}
}
