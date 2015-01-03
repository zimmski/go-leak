package leak

import (
	"runtime"
	"sync"
)

var fixRan = false
var fixMutex sync.Mutex

func fix() {
	// FIXME there is an allocation for the first started go routine, since I do not know where the allocation is coming from or how to fix that, I just want to ignore it by making an early goroutine. This does NOT show up during test execution. This is easily explained since AFAIK the tests are run in goroutines.

	fixMutex.Lock()
	defer fixMutex.Unlock()

	if fixRan {
		return
	}

	fixRan = true

	ch := func() chan bool {
		ch := make(chan bool)

		go func() {
			ch <- true

			close(ch)
		}()

		return ch
	}()

	<-ch
}

func MemoryLeaks(f func()) int {
	fix()

	var then = new(runtime.MemStats)
	var now = new(runtime.MemStats)

	runtime.GC()
	runtime.ReadMemStats(then)

	f()

	runtime.GC()
	runtime.ReadMemStats(now)

	return int((now.Mallocs - then.Mallocs) - (now.Frees - then.Frees))
}

type MemoryMark struct {
	then *runtime.MemStats
	now  *runtime.MemStats
}

func MarkMemory() *MemoryMark {
	fix()

	m := &MemoryMark{
		then: new(runtime.MemStats),
		now:  new(runtime.MemStats),
	}

	runtime.GC()
	runtime.ReadMemStats(m.then)

	return m
}

func (m *MemoryMark) Release() int {
	runtime.GC()
	runtime.ReadMemStats(m.now)

	return int((m.now.Mallocs - m.then.Mallocs) - (m.now.Frees - m.then.Frees))
}
