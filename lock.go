package main

import (
	"sync/atomic"
)

// OptimisticLock OptimisticLock
type OptimisticLock struct {
	stat int32
}

// Lock Lock
func (l *OptimisticLock) Lock() bool {
	return atomic.CompareAndSwapInt32(&l.stat, 0, 1)
}

// UnLock UnLock
func (l *OptimisticLock) UnLock() {
	l.stat = 0
}
