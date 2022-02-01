package proxy

import (
	"fmt"
	"log"
)

// video 服务对象接口
type video interface {
	playVideo(videoName string)
}

// iQiYi 爱奇艺外部连接客户端
type iQiYi struct{}

func (i *iQiYi) playVideo(videoName string) {
	fmt.Println("iQiYiL", videoName)
}

// proxyVideo 代理video对象，在服务调用之前与之后增加日志打印
type proxyVideo struct {
	cli video
}

func (p *proxyVideo) playVideo(videoName string) {
	log.Println("logging in proxy_video:", videoName)
	p.cli.playVideo(videoName)
}

func newProxyVideo() *proxyVideo {
	return &proxyVideo{cli: &iQiYi{}}
}
