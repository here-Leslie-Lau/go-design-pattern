package adapter

import (
	"strings"
	"testing"
)

func TestAdapter(t *testing.T) {
	english := newEnglish()
	talker := &languageAdapter{english: english}
	result := talker.hello()
	if strings.Contains(result, "hello") {
		t.Fail()
	}
}
