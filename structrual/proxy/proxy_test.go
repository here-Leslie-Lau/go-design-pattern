package proxy

import "testing"

func TestProxy(t *testing.T) {
	p := newProxyVideo()
	p.playVideo("春晚")
	// output:
	//2022/02/01 16:00:40 logging in proxy_video: 春晚
	//iQiYiL 春晚
}
