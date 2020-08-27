package proxy

import "testing"

func TestProxy(t *testing.T) {
	tRealSubject := RealSubject{}
	tProxy := Proxy{
		RealSubject: tRealSubject,
	}
	tProxy.Request()
}
