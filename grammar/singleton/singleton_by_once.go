package singleton

import (
	"sync"
)

var (
	// singleton *singleton
	once sync.Once
)

func Instance2() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
