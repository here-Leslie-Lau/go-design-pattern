package adapter

import "strings"

// 要适配的接口
type iEnglish interface {
	hello() string
}

func newEnglish() iEnglish {
	return &english{}
}

// english 实现
type english struct{}

func (e *english) hello() string {
	return "hello world!"
}

// talker 客户端接口
type talker interface {
	talk() string
}

// languageAdapter 语言适配器，实现客户端接口talker，翻译语言
type languageAdapter struct {
	english iEnglish
}

func (i *languageAdapter) talk() string {
	result := i.english.hello()
	if strings.Contains(result, "hello") {
		return "你好 世界！"
	}
	return result
}
