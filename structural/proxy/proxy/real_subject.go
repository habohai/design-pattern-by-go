package proxy

import "fmt"

type RealSubject struct{}

func (rs RealSubject) Request() {
	fmt.Println("real request")
}
