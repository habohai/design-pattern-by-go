package proxy

type Proxy struct {
	RealSubject
}

func (proxy Proxy) Request() {
	proxy.RealSubject.Request()
}
