package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	// 输出结果: created a mysql datasource!
	for i := 0; i < 10; i++ {
		go getDb()
	}
}
