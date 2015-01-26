package concurrency

import (
	"runtime"
	"sync/atomic"
)

const (
	UNLOCKED = 0
	LOCKED   = 1
)

type SpinLock struct {
	locked int32
}

func (sp *SpinLock) acquire() {
	while !sp.tryAcquire() { // if can't get the lock -- spin
		runtime.Gosched() // yield to other routines
	}
}

func (sp *SpinLock) tryAcquire() bool {
	return atomic.CompareAndSwapInt32((sp.locked,UNLOCKED, LOCKED))
}

func (sp *SpinLock) release() {
	sp.locked = UNLOCKED
}
