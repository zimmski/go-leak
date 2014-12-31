package leak

import (
	"runtime"
)

func MemoryLeaks(f func()) int {
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
