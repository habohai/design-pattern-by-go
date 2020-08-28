package boss

import "testing"

func TestBoss(t *testing.T) {
	//创建一个通知者
	tSubject := NewSubject("boss")
	//创建两个观察者，并与通知者绑定
	tObserverA := NewObserver("stock", "张三", tSubject)
	tObserverB := NewObserver("nba", "李四", tSubject)
	//通知者登记观察者
	tSubject.Attach(tObserverA)
	tSubject.Attach(tObserverB)
	//通知者状态改变
	tSubject.SetState("老板来了")
	//通知者开始通知观察者
	tSubject.Notify()

	secretary := NewSubject("secretary")
	tObserverA = NewObserver("stock", "王五", secretary)
	tObserverB = NewObserver("nba", "马六", secretary)
	secretary.Attach(tObserverA)
	secretary.Attach(tObserverB)
	secretary.SetState("老板来了")
	secretary.Notify()
}
