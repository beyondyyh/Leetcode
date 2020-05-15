package singleton

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

// 原子操作`atomic`配合互斥锁`mutex`可以实现非常高效的单件模式
// 增加一个原子标记位`atomic.StoreUint32(addr *uint32, val uint32)`，
// 通过原子检测标记位状态`atomic.LoadUint32(addr *unit32)`
func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}

	return instance
}
