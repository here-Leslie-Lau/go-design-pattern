package singleton

import (
	"fmt"
	"sync"
)

type datasource struct {
	name string
}

var db *datasource

var lock = &sync.Mutex{}

// 采用懒汉式单例
func getDb() *datasource {
	if db == nil {
		// 考虑并发安全
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			fmt.Println("created a mysql datasource!")
			db = &datasource{name: "mysql"}
		}
	}
	return db
}
