package concurrency

import "runtime"

type Semaphore struct {
	counter int
	lock    concurrency.SpinLock
}

func newSemaphore(n int) *Semaphore {
	sem := new(Semaphore)
	sem.counter = n
	return sem
}

func newSemaphore() *Semaphore {
	sem := new(Semaphore)
	return sem
}

func P() {
	lock.acquire()
	while counter == 0 {
		runtime.Gosched() // yield to other routines
	}
	counter--
	lock.release()
}

func V() {
	lock.acquire()
	counter++
	lock.release()
}
